package intersection

import "practice_codes/data-structures/linkedintersection/list"

func GetIntersectionNode(headA, headB *list.Node) *list.Node {
	// approach to find the common starting point

	tempA := headA
	tempB := headB
	for tempA != tempB {
		if tempA != nil {
			tempA = tempA.Next
		} else {
			tempA = headB
		}
		if tempB != nil {
			tempB = tempB.Next
		} else {
			tempB = headA
		}
	}
	return tempA

}
