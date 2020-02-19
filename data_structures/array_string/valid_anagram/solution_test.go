package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			S1 string
			S2 string
		}
		Output bool
	}{
		{struct {
			S1 string
			S2 string
		}{"anagram", "nagaram"}, true},
		{struct {
			S1 string
			S2 string
		}{"rat", "car"}, false},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.S1, tt.Input.S2)

		if got != tt.Output {
			t.Errorf("got %t, want %t", got, tt.Output)
		}
	}
}
