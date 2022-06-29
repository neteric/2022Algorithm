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
