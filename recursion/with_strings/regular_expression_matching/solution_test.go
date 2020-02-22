package solution

import "testing"

func TestSolution(t *testing.T) {
	type input struct {
		s string
		p string
	}

	tests := []struct {
		input  input
		output bool
	}{
		{input{"aa", "a"}, false},
		{input{"aa", "a*"}, true},
		{input{"ab", ".*"}, true},
		{input{"aab", "c*"}, false},
		{input{"aab", "c*a*b"}, true},
		{input{"mississippi", "mis*is*p*"}, false},
		// {input{"ab", ".*c"}, false},
		{input{"aaa", "a*a"}, true},
	}

	for _, tt := range tests {
		got := Solve(tt.input.s, tt.input.p)

		if got != tt.output {
			t.Errorf("got %t, want %t", got, tt.output)
		}
	}
}
