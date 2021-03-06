package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"abdbca", 7},
		{"cddpd", 7},
		{"pqr", 3},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		t.Skip()
		for _, tt := range tests {
			got := TopDown(tt.Input)
			if got != tt.Output {
				t.Errorf("got %d, want %d", got, tt.Output)
			}
		}
	})
}
