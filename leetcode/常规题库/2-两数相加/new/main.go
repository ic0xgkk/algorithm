package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	RL := &ListNode{}
	RS := RL
	Last := RS
	up := false
	s := 0

	for {
		if l1 == nil && l2 == nil {
			break
		}

		s = 0
		if l1 != nil {
			s += l1.Val
		}
		if l2 != nil {
			s += l2.Val
		}

		//if s >= 10 {
		//	up = true
		//	s = s % 10
		//}

		RL.Val += s

		if RL.Val >= 10 {
			up = true
			RL.Val = RL.Val % 10
		}

		Last = RL
		RL.Next = &ListNode{}
		RL = RL.Next

		if up {
			RL.Val++
			up = false
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if RL.Val == 0 {
		Last.Next = nil
	}

	return RS
}
