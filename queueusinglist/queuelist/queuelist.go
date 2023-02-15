package queuelist

import (
	"fmt"
	"strings"
)

// Node ...
type Node struct {
	data int
	next *Node
}

// Queue ...
type Queue struct {
	size  int
	front *Node
	rear  *Node
	len   int
}

// NewNode ...
func CreateNewNode(item int) *Node {
	newNode := new(Node)
	newNode.data = item
	newNode.next = nil
	return newNode
}

// InitializeQueue ...
func (q *Queue) InitializeQueue(size int) {
	q.size = size
	q.front = nil
	q.rear = nil
	q.len = 0

}

// Enque ...
func (q *Queue) Enque(item int) {
	front := q.front
	rear := q.rear
	// overflow condition
	if q.len == q.size {
		fmt.Printf("queue is overflowed, can not insert new item :: %d\n", item)
		return
	}
	newNode := CreateNewNode(item)
	// check if enqued item is first item
	if rear == nil && front == nil {
		q.rear = newNode
		q.front = newNode
		q.len++
		return
	}

	// check if queue is already having some values then make n ewnode as rear node
	rear.next = newNode
	q.rear = newNode
	q.len++

}

// Deque ...
func (q *Queue) Deque() (dquedItem int) {
	front := q.front
	rear := q.rear
	// underflow condition
	if front == nil && rear == nil {
		fmt.Printf("queue is underflow, can not deque.\n")
		return -1
	}

	// check if front and rear both are at the same location - if so, then there is only 1 item in the queue. so reset front and rear to nil
	if front == rear {
		dquedItem = front.data
		q.front = nil
		q.rear = nil
		q.len--
		return
	}
	// if dequed item is not the last item in the queue
	dquedItem = front.data
	q.front = front.next
	q.len--
	return
}

// QueuePeek ...
func (q *Queue) QueuePeek() (peek int) {
	// underflow condition
	if q.front == nil && q.rear == nil {
		fmt.Printf("queue is underflow, no peek item present.\n")
		return -1
	}
	peek = q.front.data
	return
}

// DisplayQueue ...
func (q *Queue) DisplayQueue() {
	front := q.front
	rear := q.rear
	// underflow condition
	if front == nil && rear == nil {
		fmt.Printf("queue is underflow, nothing to display.\n")
		return
	}
	QueueItems := ""
	for front != nil {
		QueueItems = QueueItems + fmt.Sprintf("%d => ", front.data)
		front = front.next
	}

	fmt.Println("Queue Items: ", strings.TrimSuffix(QueueItems, " => "))
	fmt.Printf("Front: %d && Rear: %d && queue size: %d\n", q.front.data, q.rear.data, q.len)
}
