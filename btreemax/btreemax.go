package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	maxnode := root
	maxleft := BTreeMax(root.Left)
	maxRight := BTreeMax(root.Right)
	if maxleft != nil && maxnode.Data < maxleft.Data {
		maxnode = maxleft
	}
	if maxRight != nil && maxnode.Data < maxRight.Data {
		maxnode = maxRight
	}
	return maxnode
}
