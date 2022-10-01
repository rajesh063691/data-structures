package reccalls

import "fmt"

// head recursion calls, function calls and variables are stred in stack frames
func SumOfnNumbers(n int) int {
	// define base case
	if n == 1 {
		return 1
	}
	return n + SumOfnNumbers(n-1)
}

// tail recursion, where functions calls are not stored in stack frames
func Sum(n int, result int) int {
	//base case
	if n == 1 {
		return result
	}
	return Sum(n-1, n+result)
}

func Head(n int) {
	fmt.Println("calling Head with value of n=", n)

	if n == 0 {
		return
	}
	// head recursion calls
	Head(n - 1)

	fmt.Println("returning value of n=", n)
}

func Tail(n int) {
	fmt.Println("calling Tail with value of n=", n)

	if n == 0 {
		return
	}

	fmt.Println("returning value of n=", n)
	// head recursion calls
	Tail(n - 1)
}
