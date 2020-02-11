package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			N1 int
			N2 int
		}
		Output int
	}{
		{struct {
			N1 int
			N2 int
		}{6, 9}, 3},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.N1, tt.Input.N2)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
