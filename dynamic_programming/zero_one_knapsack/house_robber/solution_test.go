package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.input)
			if got != tt.output {
				t.Errorf("got %d, want %d", got, tt.output)
			}
		}
	})
}
