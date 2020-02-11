package solution

import (
	dll "dsa/data_structures/doubly_linked_list"
	"testing"
)

func TestSolution(t *testing.T) {
	l := dll.New()
	l.PushTail(4)
	l.PushTail(3)
	l.PushTail(11)
	l.PushTail(7)

	got := Solve(l)
	want := 4

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
