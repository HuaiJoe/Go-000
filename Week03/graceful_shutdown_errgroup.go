package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Fprintln(writer, "Hello World.")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//  添加服务启动函数
	group.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Print("http server error: %v", err.Error())
		}
		return err
	})

	// 添加信号检测函数
	group.Go(Signal)

	if err := group.Wait(); err != nil {
		fmt.Println("cancel request!")
		cancel()
		fmt.Println("begin shutdown server")
		server.Shutdown(ctx)
	}
}

func Signal() error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-sigs:
		return errors.New("notify kill signals.")
	}
}
