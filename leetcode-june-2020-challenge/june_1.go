package leetcode_june_2020_challenge

/**
 * Definition for a binary tree node.
Invert a binary tree.

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil && root.Right != nil {
		root.Left = root.Right
		root.Right = nil
		invertTree(root.Left)
		return root
	} else if root.Right == nil && root.Left != nil {
		root.Right = root.Left
		root.Left = nil
		invertTree(root.Right)
		return root
	}

	RightNode := root.Right
	LeftNode := root.Left

	root.Right = LeftNode
	root.Left = RightNode

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
