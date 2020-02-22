package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type input struct {
		nums1 []int
		nums2 []int
	}

	tests := []struct {
		input  input
		output []int
	}{
		{input{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
		}, []int{2}},
		{input{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
		}, []int{9, 4}},
	}

	for _, tt := range tests {
		got := Solve(tt.input.nums1, tt.input.nums2)

		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("got %v, want %v", got, tt.output)
		}
	}
}
