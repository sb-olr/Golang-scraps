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

func (l *List) push(insert *Node) {
	curr := l.head
	for ; curr.next != nil; curr = curr.next {
	}
	if l.head == nil {
		l.head = insert
	}
	curr.next = insert
}

func (l *List) remove(del *Node) {
	for curr := l.head; curr.next != nil; curr = curr.next {
		if curr.next == del {
			curr.next = del.next
			break
		}
	}
	del = nil
}

type Node struct {
	value int
	next  *Node
}

func main() {
	a := new(List)
	c := new(Node)
	a.push(c)
	fmt.Println(a)
}
