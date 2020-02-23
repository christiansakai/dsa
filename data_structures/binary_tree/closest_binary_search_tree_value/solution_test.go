package solution

import "testing"

func TestSolution(t *testing.T) {
	tree := &node{
		val: 4,
		left: &node{
			val: 2,
			left: &node{
				val:   1,
				left:  nil,
				right: nil,
			},
			right: &node{
				val:   3,
				left:  nil,
				right: nil,
			},
		},
		right: &node{
			val:   5,
			left:  nil,
			right: nil,
		},
	}

	got := Solve(tree, 3.714286)
	if got != 4 {
		t.Errorf("got %d, want %d", got, 4)
	}
}
