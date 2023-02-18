package main

import (
	"practice_codes/data-structures/queueusingstack/queuestack"
)

func main() {
	q := queuestack.InitializeStack(5)
	q.ShowQueueItems()
	q.Enque(5)
	q.Enque(6)
	q.Enque(7)
	q.Deque()
	q.Deque()
	q.Deque()
	q.Enque(8)
	q.Enque(9)
	q.Enque(10)
	q.Deque()
	q.ShowQueueItems()

}
