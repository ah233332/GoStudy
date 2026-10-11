package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	} else {
		fmt.Println("已连接上服务端")
	}

	//reader 要在循环外只建一次：放在循环里会丢掉上一次读进缓冲区的数据
	reader := bufio.NewReader(os.Stdin) //os.stdin标准输入[终端]

	for {
		//客户端从终端读取一行用户输入，并准备发给服务端
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readerstring err=:", err)
			conn.Close()
			break
		}

		//终端输入exit,退出循环程序
		//注意：Windows 终端按回车发来的是 "\r\n"
		//导致 input 变成 "exit\r\n"，跟 "exit" 比较永远不相等。
		input := strings.Trim(line, "\r\n")
		if input == "exit" {
			fmt.Println("结束与服务端链接")
			conn.Close()
			break
		}

		//将line数据发送给服务端
		n, err := conn.Write([]byte(conn.LocalAddr().String() + ":" + line))
		if err != nil {
			fmt.Println("connwrite err=", err)
			conn.Close()
			break
		}
		fmt.Printf("客户端发送了%d字节的数据\n", n)
	}

}
