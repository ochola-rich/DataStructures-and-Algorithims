package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	min := root
	leftmin := BTreeMin(root.Left)
	rightmin := BTreeMin(root.Right)

	if leftmin != nil && min.Data > leftmin.Data {
		min = leftmin
	}
	if rightmin != nil && min.Data > rightmin.Data {
		min = rightmin
	}
	return min
}
