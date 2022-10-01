package reccalls

// calculating fibonacci numbers using head recursion
func FibHead(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fib1 := FibHead(n - 1)
	fib2 := FibHead(n - 2)
	result := fib1 + fib2
	return result
}

func FibTail(n int, a int, b int) int {
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	return FibTail(n-1, b, a+b)
}
