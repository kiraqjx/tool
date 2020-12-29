package go2

import (
	"fmt"
	"testing"
)

func TestGo2(t *testing.T) {
	first1 := &ListNode{3, nil}
	first2 := &ListNode{4, first1}
	first3 := &ListNode{2, first2}

	second1 := &ListNode{4, nil}
	second2 := &ListNode{6, second1}
	second3 := &ListNode{5, second2}

	fmt.Println(addTwoNumbers(first3, second3))
}
