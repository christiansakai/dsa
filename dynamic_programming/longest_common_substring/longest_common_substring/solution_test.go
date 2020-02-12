package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			S1 string
			S2 string
		}
		Output int
	}{
		{struct {
			S1 string
			S2 string
		}{"aab", "a"}, 1},
		{struct {
			S1 string
			S2 string
		}{"aab", "aa"}, 2},
		{struct {
			S1 string
			S2 string
		}{"aab", "baa"}, 2},
		{struct {
			S1 string
			S2 string
		}{"aab", "baab"}, 3},
		// {struct {
		// 	S1 string
		// 	S2 string
		// }{"dca", "da"}, 1},
		// {struct {
		// 	S1 string
		// 	S2 string
		// }{"abdca", "cbda"}, 2},
		// {struct {
		// 	S1 string
		// 	S2 string
		// }{"passport", "ppsspt"}, 3},
	}

	t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
		t.Skip()
		for _, tt := range tests {
			got := TopDown(tt.Input.S1, tt.Input.S2)
			if got != tt.Output {
				t.Errorf("got %d, want %d", got, tt.Output)
			}
		}
	})
}
