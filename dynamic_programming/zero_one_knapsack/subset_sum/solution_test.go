package solution

import "testing"

func TestSolution(t *testing.T) {
  tests := []struct{
    Input struct{
      Set []int
      S int
    }
    Output bool
  }{
    {
      Input: struct{
        Set []int
        S int
      }{
        Set: []int{1, 2, 3, 7},
        S: 6,
      },
      Output: true,
    },
    {
      Input: struct{
        Set []int
        S int
      }{
        Set: []int{1, 2, 7, 1, 5},
        S: 10,
      },
      Output: true,
    },
    {
      Input: struct{
        Set []int
        S int
      }{
        Set: []int{1, 3, 4, 8},
        S: 6,
      },
      Output: false,
    },
  }

  t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
    for _, tt := range tests {
    got := TopDown(tt.Input.Set, tt.Input.S)
      if got != tt.Output {
        t.Errorf("got %t, want %t", got, tt.Output)
      }
    }
  })

  // t.Run("Bottom-up Dynamic Programming", func(t *testing.T) {
  //   for _, tt := range tests {
  //   got := BottomUp(tt.Input.Set, tt.Input.S)
  //     if got != tt.Output {
  //       t.Errorf("got %t, want %t", got, tt.Output)
  //     }
  //   }
  // })

//   t.Run("Bottom-up Dynamic Programming with Minimum Space", func(t *testing.T) {
//     for _, tt := range tests {
//     got := BottomUpMinSpace(tt.Input.Set, tt.Input.S)
//       if got != tt.Output {
//         t.Errorf("got %t, want %t", got, tt.Output)
//       }
//     }
//   })
}

func TestFindSets(t *testing.T) {
  t.Skip()
}
