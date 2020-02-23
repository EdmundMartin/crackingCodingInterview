package chapter2

func KthToLast(node *Node, k int) *Node {
	count := 0
	n := node
	for n != nil {
		n = n.Next
		count++
	}
	h := node
	for i := 0; i < (count - k - 1); i++ {
		h = h.Next
	}
	return h
}