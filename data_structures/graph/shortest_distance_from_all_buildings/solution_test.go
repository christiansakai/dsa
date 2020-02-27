package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	input := [][]int{
		[]int{1, 0, 2, 0, 1},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0},
	}

	output := 7

	got := Solve(tt.input)

	if got != output {
		t.Errorf("got %d, want %d", got, output)
	}
}
