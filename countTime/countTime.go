// countTime
package main

import (
	"fmt"
	"time"
)

func count(A chan string) {
	for i := 0; i < 20; i++ {
		temp := "I can count: " + string(i) + "\n"
		fmt.Println(temp)
		A <- temp
	}
}

func main() {
	defer func(start time.Time) {
		elapsed := time.Since(start)
		fmt.Printf("function took %s \n", elapsed)
	}(time.Now())

	a := make(chan string, 20)
	go count(a)
	fmt.Println(<-a)
}
