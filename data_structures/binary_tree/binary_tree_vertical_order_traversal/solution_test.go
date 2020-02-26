package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tree1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	tree2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 0},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 7},
		},
	}

	tree3 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  9,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{
				Val:   0,
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 5},
			},
			Right: &TreeNode{Val: 7},
		},
	}

	tests := []struct {
		input  *TreeNode
		output [][]int
	}{
		{tree1, [][]int{
			[]int{9},
			[]int{3, 15},
			[]int{20},
			[]int{7},
		}},
		{tree2, [][]int{
			[]int{4},
			[]int{9},
			[]int{3, 0, 1},
			[]int{8},
			[]int{7},
		}},
		{tree3, [][]int{
			[]int{4},
			[]int{9, 5},
			[]int{3, 0, 1},
			[]int{8, 2},
			[]int{7},
		}},
	}

	for _, tt := range tests {
		got := Solve(tt.input)

		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("got %v, want %v", got, tt.output)
		}
	}
}
