func oddEvenList(head *ListNode) *ListNode {
	i := 0
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	pre := node
	pre1 := node.Next
	for node != nil {
		if i != 0 && i%2 == 0 {
			pnext := pre.Next
			nnext := node.Next
			pre.Next = node
			pre1.Next = nnext
			node.Next = pnext
			node = pre1
			pre = pre.Next
			pre1 = pre1.Next
		}
		i = i + 1
		node = node.Next
	}
	return head
}
