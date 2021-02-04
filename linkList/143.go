package linkList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	cur := slow.Next
	var pre *ListNode = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	slow.Next = nil
	left := head
	right := pre
	for right != nil {
		ln := left.Next
		rn := right.Next
		left.Next = right
		right.Next = ln
		left = ln
		right = rn
	}
}
