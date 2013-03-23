package main

import "fmt"

type list struct {
	head *element
	tail *element

	curr *element
}

func (L *list) push(E *element) {
	if L.head == nil {
		L.head = E
		L.tail = E
		E.prev = nil
		E.next = nil
	} else {
		L.tail.next = E
		E.prev = L.tail
		L.tail = E
	}
}

func (L *list) pop(E *element) {
	if L.tail != nil {
		if L.head.next == nil {
			L.head = nil
		} else {
			L.tail = L.tail.prev
			L.tail.next = nil
		}
	}
}

type element struct {
	prev *element
	next *element

	value string
}

func main() {
	fmt.Printf("Hello world\n")
	a := new(list)
	c := new(element)
	c.value = "hel"
	a.push(c)
	fmt.Printf("%5s\n", a.head.value)
}
