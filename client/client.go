package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//建立连接

	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8081")
	dial, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}
	defer dial.Close()
	for {
		inputReader := bufio.NewReader(os.Stdin)
		line, _, err := inputReader.ReadLine()
		if err != nil {
			return
		}
		str := string(line)

		nag := make([]byte, 128)
		nagNum, err := dial.Read(nag)
		if err != nil {
			return
		}
		//0=开启=false  1=关闭=true
		bol := false
		if string(nag[:nagNum]) == "0" {
			bol = false
		} else {
			bol = true
		}
		//是否开启Nagle算法 false为开启
		dial.SetNoDelay(bol)

		_, err = dial.Write([]byte(str))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		readValue := make([]byte, 128)
		read, err := dial.Read(readValue)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(readValue[:read]))
	}
}
