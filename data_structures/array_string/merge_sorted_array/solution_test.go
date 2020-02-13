package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		A1     []int
		A1Size int
		A2     []int
		A2Size int
		Output []int
	}{
		{
			A1:     []int{0, 0, 0, 0, 0, 0},
			A1Size: 0,
			A2:     []int{2, 5, 6},
			A2Size: 3,
			Output: []int{2, 5, 6, 0, 0, 0},
		},
		{
			A1:     []int{1, 2, 3, 0, 0, 0},
			A1Size: 3,
			A2:     []int{2, 5, 6},
			A2Size: 3,
			Output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			A1:     []int{4, 5, 6, 0, 0, 0},
			A1Size: 3,
			A2:     []int{1, 2, 3},
			A2Size: 3,
			Output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			A1:     []int{-1, 0, 0, 3, 3, 3, 0, 0, 0},
			A1Size: 6,
			A2:     []int{1, 2, 2},
			A2Size: 3,
			Output: []int{-1, 0, 0, 1, 2, 2, 3, 3, 3},
		},
	}

	for _, tt := range tests {
		Solve(tt.A1, tt.A1Size, tt.A2, tt.A2Size)

		if !reflect.DeepEqual(tt.A1, tt.Output) {
			t.Errorf("got %v, want %v", tt.A1, tt.Output)
		}
	}
}
