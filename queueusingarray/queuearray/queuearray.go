package queuearray

import "fmt"

// Queue ...
type Queue struct {
	size  int
	front int
	rear  int
	item  []int
}

// InitializeQueue ...
func (q *Queue) InitializeQueue(size int) {
	q.size = size
	q.front = -1
	q.rear = -1
	q.item = make([]int, size)

}

// Enque ...
func (q *Queue) Enque(item int) {
	// overflow condition
	if q.rear == q.size-1 {
		fmt.Printf("queue is overflowed, can not insert item :: %d\n", item)
		return
	}
	// check if enqued item is first item
	if q.rear == -1 && q.front == -1 {
		q.rear++
		q.front++
		q.item[q.rear] = item
		return
	}

	// check if queue is already having some values
	q.rear++
	q.item[q.rear] = item
}

// Deque ...
func (q *Queue) Deque() (dquedItem int) {
	// underflow condition
	if q.front == -1 && q.rear == -1 {
		fmt.Printf("queue is underflow, can not deque.\n")
		return -1
	}

	// check if front and rear both are at the same location - if so, then there is only 1 item in the queue. so reset front and rear to -1
	if q.front == q.rear {
		dquedItem = q.item[q.front]
		q.front = -1
		q.rear = -1
		return
	}
	// if dequed item is not the last item in the queue
	dquedItem = q.item[q.front]
	q.front++
	return
}

// QueuePeek ...
func (q *Queue) QueuePeek() (peek int) {
	// underflow condition
	if q.front == -1 && q.rear == -1 {
		fmt.Printf("queue is underflow, no peek item present.\n")
		return -1
	}
	peek = q.item[q.front]
	return
}

// DisplayQueue ...
func (q *Queue) DisplayQueue() {
	// underflow condition
	if q.front == -1 && q.rear == -1 {
		fmt.Printf("queue is underflow, nothing to display.\n")
		return
	}
	QueueItems := ""
	for i := q.front; i <= q.rear; i++ {
		QueueItems = QueueItems + fmt.Sprintf("%d ", q.item[i])
	}

	fmt.Println("Queue Items: ", QueueItems)
	fmt.Printf("Front: %d && Rear: %d\n", q.item[q.front], q.item[q.rear])
}
