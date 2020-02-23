package chapter2

// Implement an algorithm to delete a node in the middle (i.e, any node but the first and last node, not necessarily
// the exact middle) of a singly linked list given only access to that node.

func DeleteMiddleNode(node *Node) bool {
	if node == nil || node.Next == nil {
		return false
	}
	next := node.Next
	node.Next = next.Next
	node.Value = next.Value
	return true
}