// package solution

// import "testing"

// func TestSolution(t *testing.T) {
// 	tests := []struct {
// 		Input  []int
// 		Output int
// 	}{
// 		{[]int{4, 1, 2, 6, 10, 1, 12}, 32},
// 		{[]int{-4, 10, 3, 7, 15}, 25},
// 	}

// 	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
// 		t.Skip()
// 		for _, tt := range tests {
// 			got := TopDown(tt.Input)
// 			if got != tt.Output {
// 				t.Errorf("got %d, want %d", got, tt.Output)
// 			}
// 		}
// 	})
// }
