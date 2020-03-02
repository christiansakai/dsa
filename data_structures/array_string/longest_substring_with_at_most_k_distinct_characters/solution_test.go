package solution

import "testing"

func TestSolution(t *testing.T) {
	type input struct {
		s string
		k int
	}

	tests := []struct {
		input  input
		output int
	}{
		{input{"a", 1}, 1},
		{input{"abc", 1}, 1},
		{input{"abc", 2}, 2},
		{input{"abc", 3}, 3},
		{input{"eceba", 2}, 3},
		{input{"aa", 1}, 2},
	}

	for _, tt := range tests {
		got := Solve(tt.input.s, tt.input.k)
		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}
