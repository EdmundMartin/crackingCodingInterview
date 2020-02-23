package chapter2

type Node struct {
	Next *Node
	Value int
}

func NewNode(val int) *Node {
	return &Node{Value: val}
}

func (n *Node) AppendToTail(data int) {
	end := NewNode(data)
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = end
}

func DeleteNode(head *Node, data int) *Node {
	if head == nil {
		return nil
	}
	n := head
	if n.Value == data {
		return head.Next
	}
	for n.Next != nil {
		if n.Next.Value == data {
			n.Next = n.Next.Next
			return head
		}
		n = n.Next
	}
	return head
}


func SliceToLinkedList(array []int) *Node {
	if len(array) == 0 {
		return nil
	}
	head := NewNode(array[0])
	for i := 1; i < len(array); i++ {
		head.AppendToTail(array[i])
	}
	return head
}


func IterateValues(head *Node) []int {
	n := head
	found := []int{}
	for n != nil {
		found = append(found, n.Value)
		n = n.Next
	}
	return found
}