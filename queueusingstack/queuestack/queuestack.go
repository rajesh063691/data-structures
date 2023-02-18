package queuestack

import "fmt"

// Stack1 ...
type Stack1 struct {
	size int
	top  int
	item []int
}

// Stack2 ... this will be used only for swapping operation
type Stack2 struct {
	size int
	top  int
	item []int
}

// InitializeStack ...
func InitializeStack(size int) *Stack1 {
	// q1&q2 initialization
	return &Stack1{
		size: size,
		top:  -1,
		item: make([]int, size),
	}
}

// Enque ... enque will be done in Stack1
func (s1 *Stack1) Enque(item int) {
	s1.Push1(item)
}

func (s1 *Stack1) Push1(item int) {
	if s1.top == s1.size-1 {
		fmt.Printf("queue is full, can not insert item :: %d\n", item)
		return
	}
	s1.top++
	s1.item[s1.top] = item
}

func (s2 *Stack2) Push2(item int) {
	if s2.top == s2.size-1 {
		fmt.Printf("queue is overflow, can not insert item :: %d\n", item)
		return
	}
	s2.top++
	s2.item[s2.top] = item
}

// Deque ... Deque will be done in Stack1 with help of Stack2
func (s1 *Stack1) Deque() {
	s2 := &Stack2{
		size: s1.size,
		top:  -1,
		item: make([]int, s1.size),
	}

	if s1.top == -1 {
		fmt.Printf("queue is underflow, can not pop\n")
		return
	}
	// pop all the element from Stack1 and then push to Stack2
	for i := s1.top; i >= 0; i-- {
		deqItem := s1.Pop1()
		s2.Push2(deqItem)
	}
	// now pop from satck2
	deqItem := s2.Pop2()
	fmt.Printf("Item popped :: %d\n", deqItem)

	// now again pop from stack2 and Push into Stack1 to get into the original shape
	for i := s2.top; i >= 0; i-- {
		deqItem := s2.Pop2()
		s1.Push1(deqItem)
	}

}

func (s1 *Stack1) Pop1() (dequedItem int) {
	// pop from stack1
	dequedItem = s1.item[s1.top]
	s1.top--
	return
}

func (s2 *Stack2) Pop2() (dequedItem int) {
	// pop from stack2
	dequedItem = s2.item[s2.top]
	s2.top--
	return
}

func (s1 *Stack1) ShowQueueItems() {
	//var s1 *Stack1
	if s1.top == -1 {
		fmt.Println("queue is empty, nothing to display")
		return
	}

	queueItems := ""
	for i := 0; i <= s1.top; i++ {
		queueItems = queueItems + fmt.Sprintf("%d ", s1.item[i])
	}

	fmt.Println("Queue: ", queueItems)
}
