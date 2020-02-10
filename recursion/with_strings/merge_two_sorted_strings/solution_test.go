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
		}{"acu", "bst"}, "abcstu"},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.S1, tt.Input.S2)

		if got != tt.Output {
			t.Errorf("got %q, want %q", got, tt.Output)
		}
	}
}
