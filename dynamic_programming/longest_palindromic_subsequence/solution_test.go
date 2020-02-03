package solution

import "testing"

func TestSolution(t *testing.T) {
  tests := []struct{
    Input string
    Output int
  }{
    {"abdbca", 5},
    {"cddpd", 3},
    {"pqr", 1},
  }

  t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
    for _, tt := range tests {
      got := SolveTopDown(tt.Input)
      if got != tt.Output {
        t.Errorf("got %d, want %d", got, tt.Output)
      }
    }
  })

  // t.Run("Bottom-up Dynamic Programming", func(t *testing.T) {
  //   for _, tt := range tests {
  //     got := SolveBottomUp(tt.Input)
  //     if got != tt.Output {
  //       t.Errorf("got %d, want %d", got, want)
  //     }
  //   }
  // })
}
