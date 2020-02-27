package solution

import (
	"fmt"
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

	tests := []struct {
		input  *TreeNode
		output *TreeNode
	}{
		{
			input:  tree1,
			output: list1,
		},
	}

	for _, tt := range tests {
		Solve(tt.input)

		fmt.Printf(
			"%+v %+v %+v",
			tt.input,
			tt.input.Right,
			tt.input.Right.Right,
			// tt.input.Right.Right.Right,
			// tt.input.Right.Right.Right.Right,
			// tt.input.Right.Right.Right.Right.Right,
		)

		if !reflect.DeepEqual(tt.input, tt.output) {
			t.Errorf("got %v, want %v", tt.input, tt.output)
		}
	}
}
