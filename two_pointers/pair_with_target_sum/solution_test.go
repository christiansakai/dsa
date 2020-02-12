package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Arr    []int
			Target int
		}
		Output []int
	}{
		{struct {
			Arr    []int
			Target int
		}{[]int{1, 2, 3, 4, 6}, 6}, []int{1, 3}},
		{struct {
			Arr    []int
			Target int
		}{[]int{2, 5, 9, 11}, 11}, []int{0, 2}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Arr, tt.Input.Target)
		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
