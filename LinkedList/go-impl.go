package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) insertAtHead(data int) {
	newNode := &node{data: data}
	newNode.next = l.head
	l.head = newNode
	l.length++
}

func (l *linkedList) insertAtTail(data int) {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
		l.length++
		return
	}
	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = newNode
	l.length++
}

func (l *linkedList) insertAtPosition(data, position int) {
	if position > l.length || position < 0 {
		return
	}
	if position == 0 {
		l.insertAtHead(data)
		return
	}
	newNode := &node{data: data}
	currentNode := l.head
	for i := 0; i < position-1; i++ {
		currentNode = currentNode.next
	}
	newNode.next = currentNode.next
	currentNode.next = newNode
	l.length++
}

func (l *linkedList) deleteAtHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	l.length--
}

func (l *linkedList) deleteAtTail() {
	if l.head == nil {
		return
	}
	if l.head.next == nil {
		l.head = nil
		l.length--
		return
	}
	currentNode := l.head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = nil
	l.length--
}

func (l *linkedList) deleteAtPosition(position int) {
	if position > l.length || position < 0 {
		return
	}
	if position == 0 {
		l.deleteAtHead()
		return
	}
	currentNode := l.head
	for i := 0; i < position-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	l.length--
}

func (l *linkedList) print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	ll := &linkedList{}
	ll.insertAtHead(1)
	ll.insertAtHead(2)
	ll.insertAtHead(3)
	ll.insertAtTail(4)
	ll.insertAtTail(5)
	ll.insertAtTail(6)
	ll.insertAtPosition(7, 3)
	ll.insertAtPosition(8, 0)
	ll.insertAtPosition(9, 10)
	ll.deleteAtHead()
	ll.deleteAtTail()
	ll.deleteAtPosition(3)
	ll.print()
}
