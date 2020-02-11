package solution

import (
	dll "dsa/data_structures/doubly_linked_list"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	l := dll.New()
	l.PushTail(4)
	l.PushTail(3)
	l.PushTail(11)
	l.PushTail(7)

	l = Solve(l)

	got := l.ToSlice()
	want := []interface{}{7, 11, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
