package solution

func Solve(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	return p.Val == q.Val &&
		Solve(p.Left, q.Left) &&
		Solve(p.Right, q.Right)
}
