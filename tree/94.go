package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}


func inorderTraversal(root *TreeNode) []int {
	result:=[]int{}
	var traversal func(root *TreeNode) 
	traversal(root *TreeNode){
		traversal(root.Left)
		result = append(result,root.Val)
		traversal(root.Right)
	}
	traversal(root)
	return result
 }