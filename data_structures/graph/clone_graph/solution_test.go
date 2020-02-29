package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	n1 := &GraphNode{Val: 1}
	n2 := &GraphNode{Val: 2}
	n3 := &GraphNode{Val: 3}
	n4 := &GraphNode{Val: 4}

	n1.Neighbors = append(n1.Neighbors, n2)
	n2.Neighbors = append(n2.Neighbors, n3)
	n3.Neighbors = append(n3.Neighbors, n4)
	n4.Neighbors = append(n4.Neighbors, n1)

	got := Solve(n1)

	if !reflect.DeepEqual(got, n1) {
		t.Errorf("got %v, want %v", got, n1)
	}
}
