package tree

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	quene := []*TreeNode{root}
	for len(quene) > 0 {
		queneLength := len(quene)
		res := []int{}
		for i := 0; i < queneLength; i++ {
			node := quene[i]
			res = append(res, node.Val)
			if node.Left != nil {
				quene = append(quene, node.Left)
			}
			if node.Right != nil {
				quene = append(quene, node.Right)
			}
		}
		quene = quene[queneLength:]
		if len(res) > 0 {
			result = append(result, res)
		}
	}
	return result
}
