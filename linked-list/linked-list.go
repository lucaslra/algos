package linkedlist

import (
	"fmt"
)

type listNode struct {
	data interface{}
	next *listNode
	prev *listNode
}

type linkedList struct {
	head *listNode
	tail *listNode
}

func (l linkedList) display() {
	node := l.head
	for node != nil {
		fmt.Printf("%v", node.data)
		if node.next != nil {
			fmt.Print(" -> ")
		}
		node = node.next
	}

	fmt.Println()
}

func (l *linkedList) insert(data interface{}) {
	node := &listNode{next: l.head, data: data, prev: nil}
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node

	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}

	l.tail = tmp
}

func (l *linkedList) reverse() {
	node := l.head
	for node != nil {
		node.next, node.prev = node.prev, node.next
		node = node.prev
	}

	l.head, l.tail = l.tail, l.head
}

func RunDemo() {
	fmt.Println("------------------------------\nRunning Linked list demo...")
	list := linkedList{}
	list.insert("Nirvana")
	list.insert("Lucas")
	list.insert("Kate")
	list.insert("Nala")
	list.insert("Winter")
	list.insert("Summer")

	list.display()
	list.reverse()
	list.display()
	list.reverse()
	list.display()
	fmt.Print("Linked list demo completed\n------------------------------\n")
}
