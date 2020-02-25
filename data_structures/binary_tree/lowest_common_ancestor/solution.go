package solution

func Solve(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	fromLeft := Solve(root.Left, p, q)
	fromRight := Solve(root.Right, p, q)

	if fromLeft != nil && fromRight != nil {
		return root
	}

	if fromLeft != nil {
		return fromLeft
	}

	return fromRight
}
