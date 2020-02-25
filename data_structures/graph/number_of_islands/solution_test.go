package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  [][]byte
		Output int
	}{
		{[][]byte{
			[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '0'},
		}, 1},
		{[][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'},
		}, 3},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if got != tt.Output {
			t.Errorf("got %d, want %d", got, tt.Output)
		}
	}
}
