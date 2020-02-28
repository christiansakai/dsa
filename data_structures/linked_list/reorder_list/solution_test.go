package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  *ListNode
		output *ListNode
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 4},
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5},
						},
					},
				},
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 3},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		Solve(tt.input)

		if !reflect.DeepEqual(tt.input, tt.output) {
			t.Errorf("got %v, want %v", tt.input, tt.output)
		}
	}
}
