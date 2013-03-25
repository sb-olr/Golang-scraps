// writeFile
package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Scanf("%s", &filename)
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("hello")
}
