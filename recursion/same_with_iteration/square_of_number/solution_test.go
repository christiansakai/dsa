package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  int
		Output int
	}{
		{0, 0},
		{1, 1},
		{2, 4},
		{3, 9},
		{4, 16},
		{5, 25},
		{6, 36},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
