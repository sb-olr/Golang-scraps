// array
package main

import (
	"fmt"
)

func main() {
	var array1, array2 = []int{2, 3}, [2][2]string{{"a", "b"}, {"c", "d"}}
	fmt.Println(array1, array2)
}
