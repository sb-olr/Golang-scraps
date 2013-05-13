// PID
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ps", "-e", "-opid,ppid,comm")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf)
}
