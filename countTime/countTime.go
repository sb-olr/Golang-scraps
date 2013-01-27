// countTime
package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) {
		elapsed := time.Since(start)
		fmt.Printf("function took %s \n", elapsed)
	}(time.Now())
	fmt.Println("Hello, playground")
	for i := 0; i < 100; i++ {
		fmt.Printf("I can count: %d\n", i)
	}
}
