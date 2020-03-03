package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  int
		output [][]string
	}{
		{
			{4, [][]string{
				[]string{
					".Q..",
					"...Q",
					"Q...",
					"..Q.",
				},
				[]string{
					"..Q.",
					"Q...",
					"...Q",
					"..Q.",
				},
			}},
		},
	}

	for _, tt := range tests {
		got := Solve(tt.input)

		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("got %v, want %v", got, tt.output)
		}
	}
}
