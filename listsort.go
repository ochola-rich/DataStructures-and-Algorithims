package piscine

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}
	sorted := false
	for !sorted {
		sorted = true
		current := l
		for current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				sorted = false
			}
			current = current.Next
		}
	}
	return l
}
