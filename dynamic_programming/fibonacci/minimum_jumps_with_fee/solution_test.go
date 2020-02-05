package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output int
	}{
		{[]int{1, 2, 5, 2, 1, 2}, 3},
		{[]int{2, 3, 4, 5}, 5},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.Input)

			if got != tt.Output {
				t.Errorf("got %d, want %d", got, tt.Output)
			}
		}
	})
}
