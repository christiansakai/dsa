package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"abcd23145", 15},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
