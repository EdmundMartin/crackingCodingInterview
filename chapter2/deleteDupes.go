package chapter2

// Remove Dups: Write cod to remove duplicates from an unsorted linked list.
// FOLLOW UP: How would you solve this problem if no buffer is allowed?

type Set struct {
	hashMap map[int]interface{}
}

func NewSet() *Set {
	return &Set{hashMap: make(map[int]interface{})}
}

func NewSetFromSlice(array []int) *Set {
	hashMap := make(map[int]interface{})
	for _, val := range array {
		hashMap[val] = nil
	}
	return &Set{hashMap:hashMap}
}

func (s *Set) Contains(val int) bool {
	_, ok := s.hashMap[val]
	return ok
}

func (s *Set) Add(val int) {
	s.hashMap[val] = nil
}

func (s *Set) Len() int {
	return len(s.hashMap)
}


func DeleteDupesBuffer(n *Node) {
	set := NewSet()
	var prev *Node
	for n != nil {
		if set.Contains(n.Value) {
			prev.Next = n.Next
		} else {
			set.Add(n.Value)
			prev = n
		}
		n = n.Next
	}
}

