package main


type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var arr []ListNode
	for {
		if l1 == nil {
			break
		}
		arr = append(arr, *l1)
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}
	for {
		if l2 == nil {
			break
		}
		arr = append(arr, *l2)
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}




}