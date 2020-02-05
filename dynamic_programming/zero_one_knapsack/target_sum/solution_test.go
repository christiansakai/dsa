package solution

import "testing"

func TestSolution(t *testing.T) {
  tests := []struct{
    Input struct{
      Set []int
      S int
    }
    Output int
  }{
    {
      Input: struct{
        Set []int
        S int
      }{
        Set: []int{1, 1, 2, 3},
        S: 1,
      },
      Output: 3,
    },
    {
      Input: struct{
        Set []int
        S int
      }{
        Set: []int{1, 2, 7, 1},
        S: 9,
      },
      Output: 2,
    },
  }

  t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
    for _, tt := range tests {
    got := TopDown(tt.Input.Set, tt.Input.S)
      if got != tt.Output {
        t.Errorf("got %d, want %d", got, tt.Output)
      }
    }
  })
}
