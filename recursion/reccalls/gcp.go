package reccalls

// find GCP using recursion
func FindGCP(firstNum, secondNum int) int {
	if firstNum%secondNum == 0 {
		return secondNum
	}
	return FindGCP(secondNum, firstNum%secondNum)
}

//find GCP using iteration
func FindGCPIteration(firstNum, secondNum int) int {
	for secondNum > 0 {
		if firstNum%secondNum == 0 {
			return secondNum
		}
		temp := firstNum
		firstNum = secondNum
		secondNum = temp % secondNum
	}
	return 0
}
