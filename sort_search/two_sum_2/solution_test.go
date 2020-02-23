package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type input struct {
		numbers []int
		target  int
	}

	tests := []struct {
		input  input
		output []int
	}{
		{input{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{input{[]int{5, 25, 75}, 100}, []int{2, 3}},
	}

	for _, tt := range tests {
		got := Solve(tt.input.numbers, tt.input.target)

		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("got %v, want %v", got, tt.output)
		}
	}
}
