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
				[]int{},
				[]int{3},
				[]int{2},
				[]int{2, 3},
				[]int{1},
				[]int{1, 3},
				[]int{1, 2},
				[]int{1, 2, 3},
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
