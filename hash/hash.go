// hash
package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}
func main() {
	var file_n string
	fmt.Scanf("%s", &file_n)
	h1, err := getHash(file_n)
	if err != nil {
		return
	}
	fmt.Println(h1)
}
