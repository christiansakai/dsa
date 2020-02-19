package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  [][]int
		Output [][]int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},
		{
			[][]int{
				[]int{5, 1, 9, 11},
				[]int{2, 4, 8, 10},
				[]int{13, 3, 6, 7},
				[]int{15, 14, 12, 16},
			},
			[][]int{
				[]int{15, 13, 2, 5},
				[]int{14, 3, 4, 1},
				[]int{12, 6, 8, 9},
				[]int{16, 7, 10, 11},
			},
		},
	}

	for _, tt := range tests {
		Solve(tt.Input.Arr, tt.Input.K)

		if !reflect.DeepEqual(tt.Input.Arr, tt.Output) {
			t.Errorf("got %v, want %v", tt.Input.Arr, tt.Output)
		}
	}
}
