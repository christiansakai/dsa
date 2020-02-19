package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			N int
			K int
		}
		Output int
	}{
		{struct {
			N int
			K int
		}{1, 1}, 0},
		{struct {
			N int
			K int
		}{2, 1}, 0},
		{struct {
			N int
			K int
		}{2, 2}, 1},
		{struct {
			N int
			K int
		}{4, 5}, 1},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.N, tt.Input.K)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
