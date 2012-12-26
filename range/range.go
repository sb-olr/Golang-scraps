// range
package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "c", "d"}
	for u, v := range(list){
		fmt.Printf("%d: Hello World! %s\n", u, v)
	}
}
