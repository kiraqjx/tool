package go2

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	up := 0
	now := res
	for {
		isGo := false
		sum := 0
		if l1 != nil {
			isGo = true
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			isGo = true
			sum += l2.Val
			l2 = l2.Next
		}
		if up != 0 {
			isGo = true
			sum += up
		}
		if !isGo {
			break
		}
		if now.Next == nil {
			now.Next = &ListNode{}
		}
		now = now.Next
		up = sum / 10
		now.Val = sum % 10
		now.Next = nil
	}
	return res.Next
}
