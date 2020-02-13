package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output bool
	}{
		{"aba", true},
		{"abca", true},
		{"racedrivercar", false},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := SolveTopDown(tt.Input)
			if got != tt.Output {
				t.Errorf("got %t, want %t", got, tt.Output)
			}
		}
	})
}
