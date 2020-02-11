package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Arr []int
			K   int
		}
		Output int
	}{
		{struct {
			Arr []int
			K   int
		}{[]int{2, 1, 5}, 3}, 8},
		{struct {
			Arr []int
			K   int
		}{[]int{2, 1, 5, 1, 3, 2}, 3}, 9},
		{struct {
			Arr []int
			K   int
		}{[]int{2, 3, 4, 1, 5}, 2}, 7},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Arr, tt.Input.K)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
