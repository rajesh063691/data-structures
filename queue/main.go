package main

import (
	"fmt"
	"log"
	"practice_codes/data-structures/queue/queueops"
)

func main() {
	myQueue := queueops.Queue{}
	fmt.Println(myQueue)
	myQueue.Enque(1000)
	myQueue.Enque(2000)
	myQueue.Enque(3000)
	fmt.Println(myQueue)
	item, err := myQueue.DeQueue()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Item Dequeued: ", item)
	fmt.Println(myQueue)
}
