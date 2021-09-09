package breadthFirst

type TreeNode struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	quene := []*TreeNode{root}
	result := []int{}
	for len(quene) > 0 {
		length := len(quene)
		for i := 0; i < length; i++ {
			result = append(result, quene[i].Val)
			if quene[i].Left != nil {
				//quene = append(quene, quene[i].Left)
			}
			if quene[i].Right != nil {
				// quene = append(quene, quene[i].Right)
			}
		}
		quene = quene[length:]
	}
	return result
}
