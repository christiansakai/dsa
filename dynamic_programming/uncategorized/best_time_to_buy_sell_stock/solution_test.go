package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	t.Run("Bottom-up Dynamic Programming", func(t *testing.T) {
		for _, tt := range tests {
			got := BottomUp(tt.Input)

			if got != tt.Output {
				t.Errorf("got %d, want %d", got, tt.Output)
			}
		}
	})
}
