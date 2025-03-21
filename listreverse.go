package piscine

func ListReverse(l *List) {
	var current, prev, next *NodeL
	current = l.Head
	l.Tail = l.Head
	prev = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next

	}
	l.Head = prev
}
