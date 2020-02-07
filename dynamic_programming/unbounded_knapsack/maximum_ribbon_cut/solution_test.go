package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Cuttings []int
			Length   int
		}
		Output int
	}{
		{
			Input: struct {
				Cuttings []int
				Length   int
			}{
				Cuttings: []int{2, 3, 5},
				Length:   5,
			},
			Output: 2,
		},
		{
			Input: struct {
				Cuttings []int
				Length   int
			}{
				Cuttings: []int{2, 3},
				Length:   7,
			},
			Output: 3,
		},
		{
			Input: struct {
				Cuttings []int
				Length   int
			}{
				Cuttings: []int{3, 5, 7},
				Length:   13,
			},
			Output: 3,
		},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.Input.Cuttings, tt.Input.Length)

			if got != tt.Output {
				t.Errorf("got %d, want %d", got, tt.Output)
			}
		}
	})
}
