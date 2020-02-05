package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Cuttings []Cutting
			Length   int
		}
		Output int
	}{
		{
			Input: struct {
				Cuttings []Cutting
				Length   int
			}{
				Cuttings: []Cutting{
					Cutting{1, 2},
					Cutting{2, 6},
					Cutting{3, 7},
					Cutting{4, 10},
					Cutting{5, 13},
				},
				Length: 5,
			},
			Output: 14,
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
