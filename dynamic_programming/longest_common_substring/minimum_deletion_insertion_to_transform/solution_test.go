package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			S1 string
			S2 string
		}
		Output struct {
			Deletion  int
			Insertion int
		}
	}{
		{struct {
			S1 string
			S2 string
		}{"abc", "fbc"},
			struct {
				Deletion  int
				Insertion int
			}{1, 1}},
		{struct {
			S1 string
			S2 string
		}{"abdca", "cbda"},
			struct {
				Deletion  int
				Insertion int
			}{2, 1}},
		{struct {
			S1 string
			S2 string
		}{"passport", "ppsspt"},
			struct {
				Deletion  int
				Insertion int
			}{3, 1}},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.Input.S1, tt.Input.S2)
			if !reflect.DeepEqual(got, tt.Output) {
				t.Errorf("got %v, want %v", got, tt.Output)
			}
		}
	})
}
