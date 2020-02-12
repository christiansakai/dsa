package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{".,", true},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %t, want %t", got, tt.Output)
		}
	}
}
