package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	RIght *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	targetSum -= root.Val

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)

}
