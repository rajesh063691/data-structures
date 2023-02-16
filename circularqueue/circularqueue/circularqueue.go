package circularqueue

import "fmt"

// CircularQueue ...
type CircularQueue struct {
	size  int
	front int
	rear  int
	item  []int
}

// InitializeQueue ...
func (q *CircularQueue) InitializeQueue(size int) {
	q.size = size
	q.front = -1
	q.rear = -1
	q.item = make([]int, size)

}

// Enque ...
func (q *CircularQueue) Enque(item int) {
	front := q.front
	rear := q.rear

	// check if enqued item is first item
	if rear == -1 && front == -1 {
		q.rear++
		q.front++
		q.item[q.rear] = item
		return
	}
	nextRear := (rear + 1) % q.size
	// overflow condition - front is at begining and rear is at last position
	if nextRear == front {
		fmt.Printf("queue is overflowed, can not insert item :: %d\n", item)
		return
	}

	// check if queue is already having some values
	q.rear = nextRear
	q.item[q.rear] = item
}

// Deque ...
func (q *CircularQueue) Deque() (dquedItem int) {
	front := q.front
	rear := q.rear
	// underflow condition
	if front == -1 && rear == -1 {
		fmt.Printf("queue is underflow, can not deque.\n")
		return -1
	}

	// check if front and rear both are at the same location - if so, then there is only 1 item in the queue. so reset front and rear to -1
	if front == rear {
		dquedItem = q.item[front]
		q.front = -1
		q.rear = -1
		return
	}
	// deque front and increment
	dquedItem = q.item[front]
	q.front = (front + 1) % q.size
	return
}

// QueuePeek ...
func (q *CircularQueue) QueuePeek() (peek int) {
	// underflow condition
	if q.front == -1 && q.rear == -1 {
		fmt.Printf("queue is underflow, no peek item present.\n")
		return -1
	}
	peek = q.item[q.front]
	return
}

// DisplayQueue ...
func (q *CircularQueue) DisplayQueue() {
	front := q.front
	rear := q.rear
	// underflow condition
	if front == -1 && rear == -1 {
		fmt.Printf("queue is underflow, nothing to display.\n")
		return
	}
	QueueItems := ""
	for front != rear {
		QueueItems = QueueItems + fmt.Sprintf("%d ", q.item[front])
		front = (front + 1) % q.size
	}
	// finally, print the last element i.e rear element
	QueueItems = QueueItems + fmt.Sprintf("%d ", q.item[rear])

	fmt.Println("Queue Items: ", QueueItems)
	fmt.Printf("Front: %d && Rear: %d\n", q.item[q.front], q.item[q.rear])
}
