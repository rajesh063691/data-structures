package queueops

import "errors"

// Queue - struct to hold queue data

type Queue struct {
	Items []int
}

// Enque - add elements in the queue
func (q *Queue) Enque(item int) {
	q.Items = append(q.Items, item)
}

// Dequeue - removes element from first come first basis
func (q *Queue) DeQueue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("can not deque :: queue is empty")
	}
	deQueued := q.Items[0]
	q.Items = q.Items[1:]
	return deQueued, nil
}

// IsEmpty - check if queue is empty
func (q *Queue) IsEmpty() bool {
	len := len(q.Items)
	return len == 0
}

// Size - return number of elements present in queue
func (q *Queue) Size() int {
	if q.IsEmpty() {
		return 0
	}
	return len(q.Items)
}
