package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Arr []int
			N   int
		}
		Output int
	}{
		{struct {
			Arr []int
			N   int
		}{[]int{9, 8, 1, 8, 1, 7}, 5}, -1},
		{struct {
			Arr []int
			N   int
		}{[]int{9, 8, 1, 8, 1, 7}, 1}, 2},
		{struct {
			Arr []int
			N   int
		}{[]int{9, 8, 1, 8, 1, 7}, 7}, 5},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Arr, tt.Input.N)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
