package tree

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Right)
		sum += root.Val
		root.Val = sum
		traversal(root.Left)
	}
	traversal(root)
	return root
}

func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		stack := []*TreeNode{}
		for len(stack) > 0 || root != nil {
			for root != nil {
				stack = append(stack, root)
				root = root.Right
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum += root.Val
			root.Val = sum
			root = root.Left
		}
	}
	traversal(root)
	return root
}
