package tree

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return result
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		result = append(result, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return result
}
