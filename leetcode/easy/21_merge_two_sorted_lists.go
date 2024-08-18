package easy

import (
	"daily-practice/leetcode/global"
)

// mergeTwoLists uses recursion
// problem: https://leetcode.com/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	switch global.GetSolutionMode() {
	case global.SolutionModeRecursive:
		return mergeTwoListRecursive(list1, list2)
	case global.SolutionModeIterative:
		return mergeTwoListIterative(list1, list2)
	default:
		panic("solution mode missing")
	}

}

func mergeTwoListRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}

func mergeTwoListIterative(list1 *ListNode, list2 *ListNode) *ListNode {
	return nil
}

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
