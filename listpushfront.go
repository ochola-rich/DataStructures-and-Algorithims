package piscine

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: l.Head}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Head = node
	}
}
