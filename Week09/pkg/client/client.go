package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client(host string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	defer conn.Close()

	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for  {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}
		data:=string(input)
		//data := strings.TrimSpace(input)
		if data == "Q" {
			break
		}
		_, err = conn.Write([]byte(data))
		fmt.Printf("data: %s",data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}

	}
}
