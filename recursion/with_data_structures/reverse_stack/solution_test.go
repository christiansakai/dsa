package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	s := New()
	s.Push(2)
	s.Push(3)
	s.Push(5)
	s.Push(8)

	s = Solve(s)

	got := s.ToSlice()
	want := []int{8, 5, 3, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
