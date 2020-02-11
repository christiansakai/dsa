package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  int
		Output string
	}{
		{0, "0"},
		{1, "1"},
		{2, "10"},
		{3, "11"},
		{4, "100"},
		{5, "101"},
		{6, "110"},
		{11, "1011"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}
