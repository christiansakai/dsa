package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"hello", 2},
		{"world", 1},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
