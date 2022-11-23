package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	//开启监听
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go func(conn net.Conn) {
			defer conn.Close()

			for {
				_, err := conn.Write([]byte("0"))
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				var buf [128]byte
				length, err := conn.Read(buf[:])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				//0~9 48~57
				//+ 43
				index := -1
				addNumber := 0
				isExist := 0
				for i, v := range buf[:length] {
					if v == 43 {
						addNumber++
						index = i
						if addNumber > 1 {
							isExist = -1
							break
						}
					}
					if v < 48 && v > 57 && v != 43 {
						isExist = -1
						break
					}
				}
				rtn := make([]byte, 0)
				if isExist == -1 || addNumber == 0 {
					rtn = append(rtn, 48)

				} else {
					s := string(buf[:length])
					l, _ := strconv.ParseInt(s[:index], 10, 64)
					r, _ := strconv.ParseInt(s[index+1:], 10, 64)
					aa := []byte(strconv.FormatInt(l+r, 10))
					rtn = aa
				}

				_, err = conn.Write(rtn)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

			}
		}(accept)
	}
}
