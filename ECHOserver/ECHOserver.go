// ECHOserver
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("Failure to listen: %s", err.Error())
		return
	}
	for {
		if c, err := l.Accept(); (err == nil) && c != nil {
			go Echo(c)
		}
	}
}

func Echo(c net.Conn) {
	defer c.Close()
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Failure to read: %s\n", err.Error())
		return
	}
	fmt.Println("Received: ", line)
	_, err = c.Write([]byte(line))
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return
	}
}
