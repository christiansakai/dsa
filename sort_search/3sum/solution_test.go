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
		{[]int{0, 0, 0, 0}, [][]int{
			[]int{0, 0, 0},
		}},
		{[]int{1, 2, -2, -1}, [][]int{}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{
			[]int{-1, -1, 2},
			[]int{-1, 0, 1},
		}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{
			[]int{-2, 0, 2},
			[]int{-2, 1, 1},
		}},
		{[]int{3, 0, -2, -1, 1, 2}, [][]int{
			[]int{-2, -1, 3},
			[]int{-2, 0, 2},
			[]int{-1, 0, 1},
		}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
