package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Arr []int
			Key int
		}
		Output int
	}{
		{struct {
			Arr []int
			Key int
		}{[]int{1, 2, 1, 4, 5, 1}, 1}, 3},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Arr, tt.Input.Key)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
