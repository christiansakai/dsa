package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"(()", 2},
		{")()())", 4},
		{")(", 0},
		{")()(", 2},
		{"()(()", 2},
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
