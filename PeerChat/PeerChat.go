// PeerChat
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	go receive_conn()
	go send_conn()
	go receive_mess()
	go send_mess()
	select {}
}

func receive_conn() {
	c, err := net.Dial("tcp", "127.0.0.1:9999") // 9999 is port for getting port number

}

func receive_mess() {
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("Failure to listen: %s", err.Error())
		return
	}
	for {
		if c, err := l.Accept(); (err == nil) && c != nil {
			go get_mess(c)
		}
	}
}

func get_mess(c net.Conn) {
	defer c.Close()
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Failure to read: %s\n", err.Error())
		return
	}
	fmt.Println("Received: ", line)
}

func send_mess() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("Failure to dial: %s", err.Error())
		return
	}
	give_mess(c)
}

func give_mess(c net.Conn) {
	line := "Hello\n"
	_, err := c.Write([]byte(line))
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return
	}
	fmt.Println("Sent: ", line)
}
