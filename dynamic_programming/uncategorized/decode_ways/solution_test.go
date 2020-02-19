package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"", 0},
		{"1", 1},
		{"12", 2},
		{"226", 3},
		// {"0", 0},
		// {"00", 0},
		// {"10", 1},
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
