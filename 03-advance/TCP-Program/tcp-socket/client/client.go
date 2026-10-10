package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	//客户端从终端读取一行用户输入，并准备发给服务端
	reader := bufio.NewReader(os.Stdin) //os.stdin标准输入[终端]
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readerstring err=:", err)
	}

	//将line数据发送给服务端
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("connwrite err=", err)
	}
	fmt.Printf("客户端发送了%d字节的数据，并退出", n)

}
