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
		{[]int{0}, []int{0}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{0, 1}, []int{1, 0}},
		{[]int{0, 1, 0}, []int{1, 0, 0}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, tt := range tests {
		Solve(tt.Input)
		if !reflect.DeepEqual(tt.Input, tt.Output) {
			t.Errorf("got %v, want %v", tt.Input, tt.Output)
		}
	}
}
