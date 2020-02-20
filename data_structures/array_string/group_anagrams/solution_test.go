package solution

import (
	"reflect"
	"testing"
)

// This test works it just need some ordering
func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []string
		Output [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				[]string{"eat", "tea", "ate"},
				[]string{"tan", "nat"},
				[]string{"bat"},
			},
		},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
