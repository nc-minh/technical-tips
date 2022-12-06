package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

// doublyLinkedList implements a doubly linked list

type doublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *doublyLinkedList) insertAtHead(data int) {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}
	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
	l.length++
}

func (l *doublyLinkedList) insertAtTail(data int) {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
	l.length++
}

func (l *doublyLinkedList) insertAtPosition(data, position int) {
	if position > l.length || position < 0 {
		return
	}
	if position == 0 {
		l.insertAtHead(data)
		return
	}
	if position == l.length {
		l.insertAtTail(data)
		return
	}
	newNode := &node{data: data}
	currentNode := l.head
	for i := 0; i < position-1; i++ {
		currentNode = currentNode.next
	}
	newNode.next = currentNode.next
	currentNode.next.prev = newNode
	currentNode.next = newNode
	newNode.prev = currentNode
	l.length++
}

func (l *doublyLinkedList) deleteAtHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
	l.head.prev = nil
	l.length--
}

func (l *doublyLinkedList) deleteAtTail() {

	if l.head == nil {
		return
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	l.length--
}

func (l *doublyLinkedList) deleteAtPosition(position int) {

	if position > l.length || position < 0 {
		return
	}
	if position == 0 {
		l.deleteAtHead()
		return
	}
	if position == l.length {
		l.deleteAtTail()
		return
	}
	currentNode := l.head
	for i := 0; i < position-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	currentNode.next.prev = currentNode
	l.length--
}

func (l *doublyLinkedList) printList() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	dll := &doublyLinkedList{}
	dll.insertAtHead(1)
	dll.insertAtHead(2)
	dll.insertAtHead(3)
	dll.insertAtHead(4)
	dll.insertAtHead(5)
	dll.insertAtTail(6)
	dll.insertAtTail(7)
	dll.insertAtTail(8)
	dll.insertAtTail(9)
	dll.insertAtTail(10)
	dll.insertAtPosition(11, 5)
	dll.insertAtPosition(12, 10)
	dll.insertAtPosition(13, 15)
	dll.deleteAtHead()
	dll.deleteAtTail()
	dll.deleteAtPosition(5)
	dll.printList()
}
