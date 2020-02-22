package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		Input struct {
			S    string
			Dict []string
		}
		Output bool
	}{
		{struct {
			S    string
			Dict []string
		}{"leetcode", []string{"leet", "code"}}, true},
		// {struct{
		//   S string
		//   Dict []string
		// }{"applepenapple", []string{"apple", "pen"}}, true},
		// {struct{
		//   S string
		//   Dict []string
		// }{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
	}

	t.Run("Top-down Dynamic Programming", func(t *testing.T) {
		for _, tt := range tests {
			got := TopDown(tt.Input.S, tt.Input.Dict)

			if got != tt.Output {
				t.Errorf("got %t, want %t", got, tt.Output)
			}
		}
	})
}
