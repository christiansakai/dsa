package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			Nums   []int
			Target int
		}
		Output []int
	}{
		{struct {
			Nums   []int
			Target int
		}{[]int{}, 0}, []int{-1, -1}},
		{struct {
			Nums   []int
			Target int
		}{[]int{1}, 1}, []int{0, 0}},
		{struct {
			Nums   []int
			Target int
		}{[]int{2, 2}, 2}, []int{0, 1}},
		{struct {
			Nums   []int
			Target int
		}{[]int{1, 1, 1, 5}, 1}, []int{0, 2}},
		{struct {
			Nums   []int
			Target int
		}{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{struct {
			Nums   []int
			Target int
		}{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
	}

	for _, tt := range tests {
		got := Solve(tt.Input.Nums, tt.Input.Target)
		if !reflect.DeepEqual(got, tt.Output) {
			t.Errorf("got %v, want %v", got, tt.Output)
		}
	}
}
