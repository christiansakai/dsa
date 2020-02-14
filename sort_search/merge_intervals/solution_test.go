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
		{[][]int{
			[]int{1, 3},
			[]int{2, 6},
			[]int{8, 10},
			[]int{15, 18},
		}, [][]int{
			[]int{1, 6},
			[]int{8, 10},
			[]int{15, 18},
		}},
		{[][]int{
			[]int{1, 4},
			[]int{4, 5},
		}, [][]int{
			[]int{1, 5},
		}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
