package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{"the sky is blue", "blue is sky the"},
		// {"  hello world!  ", "world! hello"},
		// {"a good   example", "example good a"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}
