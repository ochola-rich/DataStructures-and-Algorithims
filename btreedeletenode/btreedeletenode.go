package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}
	if node.Left == nil && node.Right == nil {
		return BTreeTransplant(root, node, nil)
	}

	if node.Right == nil {
		return BTreeTransplant(root, node, root.Left)
	}

	suc := BTreeMin(node.Right)

	if suc.Parent != node {
		root = BTreeTransplant(root, suc, suc.Right)
		suc.Right = node.Right
		if suc.Right != nil {
			suc.Right.Parent = suc
		}
	}
	root = BTreeTransplant(root, node, suc)
	suc.Left = node.Left
	if suc.Left != nil {
		suc.Left.Parent = suc
	}
	return root
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}

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
  func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem == root.Data {
		return root
	} else {
		return root
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right != nil && root.Right.Data < root.Data || root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	return BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			root.Left = BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			root.Right = BTreeInsertData(root.Right, data)
		}
	}
	return root
}
