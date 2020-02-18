package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var l1 *ListNode = nil
	var l1want *ListNode = nil

	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2want := &ListNode{
		Val:  1,
		Next: nil,
	}

	l3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	l3want := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	l4 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l4want := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l5 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	l5want := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	tests := []struct {
		Input  *ListNode
		Output *ListNode
	}{
		{l1, l1want},
		{l2, l2want},
		{l3, l3want},
		{l4, l4want},
		{l5, l5want},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
