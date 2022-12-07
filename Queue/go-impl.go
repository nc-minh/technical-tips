package main

import (
	"sync"
)

type QueueADT struct {
	sync.Mutex
	items []int
	size  int
}

func (q *QueueADT) Enqueue(item int) {
	q.Lock()
	defer q.Unlock()
	q.items = append(q.items, item)
	q.size++
}

func (q *QueueADT) Dequeue() int {
	q.Lock()
	defer q.Unlock()
	if q.size == 0 {
		return 0
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.size--
	return item
}

func (q *QueueADT) Size() int {
	return q.size
}

func (q *QueueADT) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueADT) Peek() int {
	if q.size == 0 {
		return 0
	}
	return q.items[0]
}

func main() {
	var q QueueADT
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)

	println(q.Peek())
	q.Dequeue()
	println(q.Peek())

	println(q.Size())

}
