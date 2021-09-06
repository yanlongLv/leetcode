package tree

// * Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	var pre *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root == pre {
			result = append(result, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return result
}

func postorderTraversal1(root *TreeNode) []int {
	result := []int{}
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		result = append(result, root.Val)
	}
	traversal(root)
	return result
}
