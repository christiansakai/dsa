package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  int
		Output int
	}{
		{4, 2},
		{8, 2},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
