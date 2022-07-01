package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkListBySlice(in []int) *ListNode {
	if len(in) <= 0 {
		return nil
	}
	var head = new(ListNode)
	head.Val = in[0]

	p := head

	for i := 1; i < len(in); i++ {
		newNode := new(ListNode)
		newNode.Val = in[i]
		p.Next = newNode
		p = p.Next
	}

	return head
}

func traverseList(head *ListNode) []int {

	var r []int

	for head != nil {
		r = append(r, head.Val)
		head = head.Next
	}
	return r

}

// number: 21
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := new(ListNode)
	p := dummy

	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	} else {
		p.Next = p2
	}

	return dummy.Next

}
