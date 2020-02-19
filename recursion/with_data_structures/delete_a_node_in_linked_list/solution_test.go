package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	l1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	l1Want := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	Solve(l1.Next)

	if !reflect.DeepEqual(l1, l1Want) {
		t.Errorf("got %v, want %v", l1, l1Want)
	}

	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	l2Want := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	Solve(l2.Next.Next)

	if !reflect.DeepEqual(l2, l2Want) {
		t.Errorf("got %v, want %v", l2, l2Want)
	}
}
