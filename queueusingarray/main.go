package main

import (
	"fmt"
	"practice_codes/data-structures/queueusingarray/queuearray"
)

func main() {
	q := queuearray.Queue{}
	q.InitializeQueue(5)
	q.DisplayQueue()
	q.Deque()
	q.Enque(5)
	//q.DisplayQueue()
	q.Enque(6)
	q.Enque(7)
	q.Enque(8)
	q.Enque(9)
	// q.Enque(10)
	// q.Enque(61)
	// q.Enque(62)
	// q.Enque(63)
	q.DisplayQueue()
	fmt.Printf("Dequed item :: %d\n", q.Deque())
	fmt.Printf("Dequed item :: %d\n", q.Deque())
	q.DisplayQueue()
	fmt.Printf("Peek item :: %d\n", q.QueuePeek())
}
