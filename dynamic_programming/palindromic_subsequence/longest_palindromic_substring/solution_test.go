package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{"abdbca", "bdb"},
		{"cddpd", "dpd"},
		{"pqr", "p"},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.Input)
			if got != tt.Output {
				t.Errorf("got %q, want %q", got, tt.Output)
			}
		}
	})
}
