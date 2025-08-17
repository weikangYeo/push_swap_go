package model

type Node struct {
	Value    int
	NextNode *Node
}

// find index of the nodeValue in current Node list, -1 if not found
func (head *Node) IndexOf(nodeValue int) int {

	i := 0
	ptr := head

	for ptr != nil {
		if ptr.Value == nodeValue {
			return i
		}
		ptr = ptr.NextNode
		i++
	}
	return -1
}

func (head *Node) Length() int {

	i := 0
	ptr := head

	for ptr != nil {
		ptr = ptr.NextNode
		i++
	}
	return i
}

func (head *Node) GetArrValue() []int {
	size := head.Length()
	if size == 0 {
		return []int{}
	}

	arr := make([]int, size)
	ptr := head
	for ptr != nil {
		arr = append(arr, ptr.Value)
		ptr = ptr.NextNode
	}
	return arr
}

func (head *Node) IsSorted() bool {
	ptr := head
	if ptr == nil {
		return false
	}

	previousValue := ptr.Value
	ptr = ptr.NextNode

	for ptr != nil {
		if ptr.Value < previousValue {
			return false
		}
		ptr = ptr.NextNode
	}
	return true
}
