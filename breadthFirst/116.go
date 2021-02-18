package breadthFirst

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	quene := []*Node{root}
	for len(quene) > 0 {
		endLength := len(quene)
		for i := 0; i < endLength; i++ {
			if i+1 >= endLength {
				quene[i].Next = nil
			} else {
				quene[i].Next = quene[i+1]
			}
			if quene[i].Left != nil {
				quene = append(quene, quene[i].Left)
			}
			if quene[i].Right != nil {
				quene = append(quene, quene[i].Right)
			}
		}
		quene = quene[endLength:]
	}
	return root
}

func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	// 返回根节点
	return root
}
