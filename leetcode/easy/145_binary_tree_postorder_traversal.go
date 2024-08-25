package easy

import "daily-practice/models"

// postorderTraversal uses recursive
// problem: https://leetcode.com/problems/binary-tree-postorder-traversal/description/
func postorderTraversal(root *models.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
}
