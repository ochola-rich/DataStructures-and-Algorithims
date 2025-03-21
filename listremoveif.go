package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head.Data == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	var prev *NodeL
	current := l.Head
	for current != nil {
		if data_ref == current.Data {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}
	l.Tail = prev
}
