package reccalls

// implement using head recursion
func FactorialHead(n int) int {
	// base case
	if n == 1 {
		return 1
	}
	result1 := FactorialHead(n - 1)
	result := n * result1
	return result
}

// Factorial using tail recursion
func FactorialTail(n int, accumalator int) int {
	if n == 1 {
		return accumalator
	}
	return FactorialTail(n-1, n*accumalator)
}
