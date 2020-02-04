package solution

import "testing"

func TestSolution(t *testing.T) {
  tests := []struct{
    Input []int
    Output int
  }{
    {[]int{1, 2, 3, 9}, 3},
    {[]int{1, 2, 7, 1, 5}, 0},
    {[]int{1, 3, 100, 4}, 92},
  }

  t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
    for _, tt := range tests {
    got := TopDown(tt.Input)
      if got != tt.Output {
        t.Errorf("got %d, want %d", got, tt.Output)
      }
    }
  })

  // t.Run("Bottom-up Dynamic Programming", func(t *testing.T) {
  //   for _, tt := range tests {
  //   got := BottomUp(tt.Input)
  //     if got != tt.Output {
  //       t.Errorf("got %t, want %t", got, tt.Output)
  //     }
  //   }
  // })

  // t.Run("Bottom-up Dynamic Programming with Minimum Space", func(t *testing.T) {
  //   for _, tt := range tests {
  //   got := BottomUpMinSpace(tt.Input)
  //     if got != tt.Output {
  //       t.Errorf("got %t, want %t", got, tt.Output)
  //     }
  //   }
  // })
}

func TestFindSets(t *testing.T) {
  t.Skip()
}
