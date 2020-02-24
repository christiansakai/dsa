package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
