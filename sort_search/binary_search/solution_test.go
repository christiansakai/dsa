package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Nums   []int
			Target int
		}
		Output int
	}{
		{struct {
			Nums   []int
			Target int
		}{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{struct {
			Nums   []int
			Target int
		}{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Nums, tt.Input.Target)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
