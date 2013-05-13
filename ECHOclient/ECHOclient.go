// ECHOclient
package main

import (
	"fmt"
	"net"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("Failure to dial: %s", err.Error())
		return
	}
	Send(c)
}

func Send(c net.Conn) {
	line := "Hello\n"
	_, err := c.Write([]byte(line))
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return
	}
	fmt.Println("Sent: ", line)
}
