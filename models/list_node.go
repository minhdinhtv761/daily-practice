package models

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	node := &ListNode{}
	node.Val = vals[0]
	if len(vals) > 1 {
		node.Next = NewListNode(vals[1:])
	}
	return node
}

func (node *ListNode) GetVals() []int {
	if node == nil {
		return []int{}
	}
	return append([]int{node.Val}, node.Next.GetVals()...)
}
