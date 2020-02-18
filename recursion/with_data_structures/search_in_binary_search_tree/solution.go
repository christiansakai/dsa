package solution

func Solve(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	} else if val < root.Val {
		return Solve(root.Left, val)
	} else {
		return Solve(root.Right, val)
	}
}
