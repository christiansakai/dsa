package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	g := New()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	got := Solve(g)
	want := []int{0, 1, 2, 4, 3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
