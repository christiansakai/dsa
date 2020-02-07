package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Coins []int
			Total int
		}
		Output int
	}{
		{
			Input: struct {
				Coins []int
				Total int
			}{
				Coins: []int{1, 2, 3},
				Total: 5,
			},
			Output: 5,
		},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.Input.Coins, tt.Input.Total)

			if got != tt.Output {
				t.Errorf("got %d, want %d", got, tt.Output)
			}
		}
	})
}
