package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{"hello world", "helloworld"},
		{"hello\tworld", "helloworld"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}
