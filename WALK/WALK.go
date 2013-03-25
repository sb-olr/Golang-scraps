// WALK
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("C:\\Thing_Thing", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
