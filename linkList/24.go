package linkList

// //  Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func swapPairs(head *ListNode) *ListNode {
	dumny := &ListNode{Next: head}
	current := dumny
	for current.Next != nil && current.Next.Next != nil {
		temp1 := current.Next
		temp2 := current.Next.Next
		current.Next = temp2
		temp1.Next = temp2.Next
		temp2.Next = temp1
		current = temp1
	}
	return dumny.Next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs1(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairs2(head *ListNode) *ListNode {
	dumny := &ListNode{Next: head}
	current := dumny
	for current.Next != nil && current.Next.Next != nil {
		temp1 := current.Next
		temp2 := current.Next.Next
		current.Next = temp2
		temp1.Next = temp2.Next
		temp2.Next = temp1
		current = temp1

	}
	return dumny.Next
}
