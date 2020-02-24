package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"I", 1},
		{"IV", 4},
		{"V", 5},
		{"IX", 9},
		{"X", 10},
		{"XX", 20},
		{"XL", 40},
		{"L", 50},
		{"XC", 90},
		{"C", 100},
		{"CD", 400},
		{"D", 500},
		{"CM", 900},
		{"M", 1000},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
