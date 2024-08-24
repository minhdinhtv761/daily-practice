package easy

import "daily-practice/models"

// mergeTwoLists uses recursion
// problem: https://leetcode.com/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *models.ListNode, list2 *models.ListNode) *models.ListNode {
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
