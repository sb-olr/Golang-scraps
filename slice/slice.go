// slice
package main

import (
	"fmt"
)

func main() {
	var slice1 = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	slice1 = append(slice1, 0, -1)
	fmt.Println(slice1)
}