package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input  []int
		Output []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)

		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
