package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"", 0},
		{"  ", 0},
		{"  -", 0},
		{"42", 42},
		{"  -42", -42},
		{"+1", 1},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"3.14159", 3},
		{"10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459", 2147483647},
		{"  0000000000012345678", 12345678},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
