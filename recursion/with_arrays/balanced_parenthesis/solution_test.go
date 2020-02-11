package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output bool
	}{
		{"()", true},
		{"(())", true},
		{"(())()", true},
		{"((()))((()))", true},
		{"(", false},
		{")()(", false},
		{"(()()()()", false},
		{"((())))((((()", false},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %t, want %t", got, tt.Output)
		}
	}
}
