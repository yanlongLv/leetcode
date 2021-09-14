package linkList

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

// 你应当 保留 两个分区中每个节点的初始相对位置。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/partition-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func partition(head *ListNode, x int) *ListNode {
	node1 := &ListNode{}
	cur1 := node1
	node2 := &ListNode{}
	cur2 := node2
	for head != nil {
		if head.Val < x {
			node1.Next = head
			node1 = node1.Next

		} else {
			node2.Next = head
			node2 = node2.Next
		}
		head = head.Next

	}
	node2.Next = nil

	node1.Next = cur2.Next
	return cur1.Next
}
