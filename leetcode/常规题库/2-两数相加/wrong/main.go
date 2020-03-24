package wrong

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var L1 []int
	var L2 []int
	var R []int
	for {
		if l1 == nil && l2 == nil {
			break
		}
		L1 = append(L1, l1.Val)
		L2 = append(L2, l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}

	up := false
	l := len(L1)
	for i:=0; i<l; i++ {
		s := 0
		if up == true {
			s =  L1[i] + L2[i] + 1
			up = false
		} else
		{
			s =  L1[i] + L2[i]
		}

		if s >= 10 {
			s = s % 10
			up = true
		}
		R = append(R, s)
	}
	if up == true {
		R = append(R, 1)
	}

	r := &ListNode{}
	rr := r

	L := len(R)

	for i := 0; i < L; i++ {
		r.Val = R[i]
		if i == L - 1 {
			break
		}
		r.Next = &ListNode{}
		r = r.Next
	}
	return rr
}
