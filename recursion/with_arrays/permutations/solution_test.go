package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
