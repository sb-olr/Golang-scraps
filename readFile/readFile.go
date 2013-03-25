// readFile
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("C:/Users/woW/Desktop/GitHubLog.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[0:n])
	}

	// another way to read file
	buf2, err := ioutil.ReadFile("C:/Users/woW/Desktop/GitHubLog.txt")
	if err != nil {
		return
	}
	fmt.Println(buf2)
}
