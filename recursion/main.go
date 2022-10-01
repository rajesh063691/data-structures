package main

import (
	"fmt"
	"practice_codes/data-structures/recursion/reccalls"
)

func main() {
	// sum := reccalls.SumOfnNumbers(5)
	// fmt.Println("Head Recusrion sum: ", sum)
	// sum = reccalls.Sum(5, 1)
	// fmt.Println("Tail Recusrion sum: ", sum)

	// reccalls.Head(5)
	// reccalls.Tail(5)

	// fmt.Println("Factorial of 5 using head recusrion: ", reccalls.FactorialHead(5))
	// fmt.Println("Factorial of 5 using tail recusrion: ", reccalls.FactorialTail(5, 1))

	// fmt.Println("Fibonacci of 5 using head: ", reccalls.FibHead(5))
	// fmt.Println("Fibonacci of 5 using Tail: ", reccalls.FibTail(10, 0, 1))

	fmt.Printf("GCP of 15 and 5 using recursion = %d\n", reccalls.FindGCP(15, 5))
	fmt.Printf("GCP of 24 and 16  using recursion = %d\n", reccalls.FindGCP(24, 16))

	fmt.Printf("GCP of 24 and 16 using iteration = %d\n", reccalls.FindGCPIteration(15, 5))
	fmt.Printf("GCP of 24 and 16 using iteration = %d\n", reccalls.FindGCPIteration(24, 16))

}
