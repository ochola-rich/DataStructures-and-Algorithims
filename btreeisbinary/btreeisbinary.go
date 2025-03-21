package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right != nil && root.Right.Data < root.Data || root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	return BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right)
}
