package piscine

// Write a function ListForEach that applies a function given as argument to the data within each node of the list l.
// The function given as argument must have a pointer as argument: l *List
// Copy the functions Add2_node and Subtract3_node in the same file as the function ListForEach is defined.

func ListForEach(l *List, f func(*NodeL)) {
	if l == nil {
		return
	}
	node := l.Head
	for node != nil {
		f(node)
		node = node.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
