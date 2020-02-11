package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Arr []int
			S   int
		}
		Output int
	}{
		{struct {
			Arr []int
			S   int
		}{[]int{2, 1, 5, 2, 3, 2}, 7}, 2},
		{struct {
			Arr []int
			S   int
		}{[]int{2, 1, 5, 2, 8}, 7}, 1},
		{struct {
			Arr []int
			S   int
		}{[]int{3, 4, 1, 1, 6}, 8}, 3},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Arr, tt.Input.S)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
