package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			S1 string
			S2 string
		}
		Output string
	}{
		{struct {
			S1 string
			S2 string
		}{"0", "0"}, "0"},
		{struct {
			S1 string
			S2 string
		}{"0", "1"}, "1"},
		{struct {
			S1 string
			S2 string
		}{"1", "0"}, "1"},
		{struct {
			S1 string
			S2 string
		}{"1", "1"}, "10"},
		{struct {
			S1 string
			S2 string
		}{"10", "1"}, "11"},
		{struct {
			S1 string
			S2 string
		}{"10", "10"}, "100"},
		{struct {
			S1 string
			S2 string
		}{"10", "11"}, "101"},
		{struct {
			S1 string
			S2 string
		}{"11", "11"}, "110"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.S1, tt.Input.S2)
		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}
