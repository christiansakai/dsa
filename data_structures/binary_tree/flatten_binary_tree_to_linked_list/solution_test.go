package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}

	list1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	tree2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}

	list2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}
	tests := []struct {
		input  *TreeNode
		output *TreeNode
	}{
		{
			input:  tree1,
			output: list1,
		},
		{
			input:  tree2,
			output: list2,
		},
	}

	for _, tt := range tests {
		Solve(tt.input)

		if !reflect.DeepEqual(tt.input, tt.output) {
			t.Errorf("got %v, want %v", tt.input, tt.output)
		}
	}
}
