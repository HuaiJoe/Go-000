package server

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type TCPServer struct {
	Port  int
	Host  string
	Close chan struct{}
}

func NewTCPServer(host string, port int) *TCPServer {
	return &TCPServer{
		Port:  port,
		Host:  host,
		Close: make(chan struct{}),
	}
}
func (s *TCPServer) Start() {
	if s.Port == 0 {
		s.Port = 8080
	}
	if s.Host == "" {
		s.Host = "localhost"
	}
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		log.Fatalf("start server failed.detail info: %v \n", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cons := make(chan net.Conn, 10)

	// 接受请求
	go accept(ctx, listen, cons)
	// 2个处理请求
	for i := 0; i < 2; i++ {
		go serve(ctx, s.Close, cons)
	}
	// 退出程序
	select {
	case <-s.Close:
		cancel()
		// 不再接收请求
		close(cons)
	}
	// 1后关闭监听，让在途请求有10S的处理时间
	<-time.After(1 * time.Second)
	listen.Close()
	fmt.Print("close server")
}

func accept(ctx context.Context, listen net.Listener, cons chan<- net.Conn) {
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf(" accept connection error.detail info: %v \n", err)
			continue
		}
		select {
		case err := <-ctx.Done():
			log.Printf("ctx error.detail info: %v \n", err)
			return
		case cons <- conn:
			log.Print("create connection succeed")
		}
	}
}
func serve(ctx context.Context, close chan<- struct{}, cons <-chan net.Conn) {
	for {
		select {
		case err := <-ctx.Done():
			log.Fatalf("ctx failed.detail info: %v \n", err)

			return
		case conn := <-cons:
			handle(conn, close)
		}
	}
}

func handle(conn net.Conn, close chan<- struct{}) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	data := make([]byte, 0)
	for {
		body, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Println("read  request data finished")
			} else {
				fmt.Printf("read  request data failed. %v", err)
			}
			break
		}
		// 请求体接收结束
		if string(body) == "$" {
			break
		}
		data = append(data, body...)
	}
	if len(data) == 0 {
		writer.WriteString("receive data error")
	}
	// 接收到推出命令后，开始准备关闭服务器
	if strings.EqualFold(string(data), "quit") {
		close <- struct{}{}
	} else {
		writer.WriteString(doBuz(data))
	}
	writer.Flush()
	err := conn.Close()
	if err != nil {
		fmt.Printf("close conn error : %v",err)
	}
}

func doBuz(data []byte) string {
	fmt.Printf("request data: %s", data)
	return "succeed"
}
