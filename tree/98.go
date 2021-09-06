package tree

import "math"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isValidBST(root *TreeNode) bool {
	var minInt int = math.MinInt64
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= minInt {
			return false
		}
		minInt = root.Val
		root = root.Right
	}
	return true
}
