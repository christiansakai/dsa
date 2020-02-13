package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		L1     *ListNode
		L2     *ListNode
		Output *ListNode
	}{
		{
			L1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			L2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			Output: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			L1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
			L2: &ListNode{
				Val:  9,
				Next: nil,
			},
			Output: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		got := Solve(tt.L1, tt.L2)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
