package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	l7 := &ListNode{Val: 7}
	l13 := &ListNode{Val: 13}
	l11 := &ListNode{Val: 11}
	l10 := &ListNode{Val: 10}
	l1 := &ListNode{Val: 1}

	l7.Next = l13
	l13.Next = l11
	l11.Next = l10
	l10.Next = l1

	l13.Random = l7
	l11.Random = l1
	l10.Random = l11
	l1.Random = l7

	got := Solve(l7)

	if !reflect.DeepEqual(got, l7) {
		t.Errorf("got %v, want %v", got, l7)
	}
}
