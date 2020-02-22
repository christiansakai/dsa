package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  [][]int
		Output []int
	}{
		// {
		//   [][]int{
		//     []int{6, 9, 7},
		//   },
		//   []int{6, 9, 7},
		// },
		// {
		//   [][]int{
		//     []int{3},
		//     []int{2},
		//   },
		//   []int{3, 2},
		// },
		// {
		//   [][]int{
		//     []int{7},
		//     []int{9},
		//     []int{6},
		//   },
		//   []int{7, 9, 6},
		// },
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
