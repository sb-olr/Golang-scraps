// LinkedList
package main

import (
	"fmt"
)

type List struct {
	head   *Node
	length int
}

func (l *List) insertBefore(n *Node, insert *Node) {
	if l.head == n {
		l.head = insert
	}
	insert.next = n
	l.length++
}

func (l *List) insertAfter(n *Node, insert *Node) {
	insert.next = n.next
	n.next = insert
	l.length++
}

func (l *List) remove()

type Node struct {
	value int
	next  *Node
}

func main() {

}
