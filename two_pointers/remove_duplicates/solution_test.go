package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Skip()
	tests := []struct {
		Input  []int
		Output struct {
			Arr []int
			Len int
		}
	}{
		{[]int{2, 3, 3, 3, 6, 9, 9}, struct {
			Arr []int
			Len int
		}{[]int{2, 3, 6, 9, 0, 0, 0}, 4}},
		{[]int{2, 2, 2, 11}, struct {
			Arr []int
			Len int
		}{[]int{2, 11, 0, 0}, 2}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input)
		if !reflect.DeepEqual(tt.Input, tt.Output.Arr) {
			t.Errorf("got %v, want %v", tt.Input, tt.Output.Arr)
		}

		if got != tt.Output.Len {
			t.Errorf("got %d, want %d", got, tt.Output.Len)
		}
	}
}
