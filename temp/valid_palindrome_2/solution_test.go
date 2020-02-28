package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"aba", true},
		{"abca", true},
		{"racecar", true},
		{"racbecar", true},
		{"racbexcar", false},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.input)
			if got != tt.output {
				t.Errorf("got %t, want %t", got, tt.output)
			}
		}
	})
}
