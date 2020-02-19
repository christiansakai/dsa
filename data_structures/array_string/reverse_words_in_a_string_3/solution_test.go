package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{"", ""},
		{"a", "a"},
		{"an", "na"},
		{"ant", "tna"},
		{"ant man", "tna nam"},
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %s, want %s", got, tt.Output)
		}
	}
}
