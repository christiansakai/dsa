package solution

import "testing"

func TestSolution(t *testing.T) {
	l := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	got := Solve(l)
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
