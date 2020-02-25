package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{5, 1, 1}, []int{1, 1, 5}},
	}

	for _, tt := range tests {
		Solve(tt.Input)

		if !reflect.DeepEqual(tt.Input, tt.Output) {
			t.Errorf("got %v, want %v", tt.Input, tt.Output)
		}
	}
}
