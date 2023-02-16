package main

import (
	"practice_codes/data-structures/circularqueue/circularqueue"
)

func main() {
	q := circularqueue.CircularQueue{}
	q.InitializeQueue(5)
	q.DisplayQueue()
	q.Enque(1)
	q.Enque(2)
	q.Enque(3)
	q.Enque(4)
	q.Enque(5)
	q.Deque()
	q.Enque(6)
	q.Deque()
	q.Deque()
	q.Deque()
	// q.Enque(6)
	// q.Enque(7)
	q.DisplayQueue()

}
