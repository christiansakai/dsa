package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			X float64
			N int
		}
		Output float64
	}{
		{struct {
			X float64
			N int
		}{2.0000, 10}, 1024.0000},
		{struct {
			X float64
			N int
		}{2.0000, -2}, 0.25000},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.X, tt.Input.N)

		if got != tt.Output {
			t.Errorf("got %f, want %f", got, tt.Output)
		}
	}
}
