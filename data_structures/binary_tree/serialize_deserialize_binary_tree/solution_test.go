package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var t1 *TreeNode = nil

	t2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	t3 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	tests := []struct {
		Data *TreeNode
	}{
		{t1},
		{t2},
		{t3},
	}

	for _, tt := range tests {
		serialized := SolveSerialize(tt.Data)
		got := SolveDeserialize(serialized)

		if !reflect.DeepEqual(got, tt.Data) {
			t.Errorf("got %v, want %v", got, tt.Data)
		}
	}
}
