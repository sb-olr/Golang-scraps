// factorial
package main

import (
	"fmt"
)

func main() {
	f := func(i int) int{
		var count = 1
		for j := 1; j <= i; j++{
			count *= j
		}
		return count
	}
	fmt.Println(f(9))
}
