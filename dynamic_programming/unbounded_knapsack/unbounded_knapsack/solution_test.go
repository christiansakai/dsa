package solution

import "testing"

func TestSolution(t *testing.T) {
  tests := []struct{
    Input struct{
      Fruits []Fruit
      Capacity int
    }
    Output int
  }{
    {
      Input: struct {
        Fruits []Fruit
        Capacity int
      }{
        Fruits: []Fruit{
          Fruit{"Apple", 1, 15},
          Fruit{"Orange", 2, 20},
          Fruit{"Melon", 3, 50},
        },
        Capacity: 5,
      },
      Output: 80,
    },
  }

  t.Run("Top-down Dynamic Programming with Memoization", func(t *testing.T) {
    for _, tt := range tests {
      got := TopDown(tt.Input.Fruits, tt.Input.Capacity)

      if got != tt.Output {
        t.Errorf("got %d, want %d", got, tt.Output)
      }
    }
  })

  t.Run("Bottom-up Dynamic Programming", func(t *testing.T) {
    for _, tt := range tests {
      got := BottomUp(tt.Input.Fruits, tt.Input.Capacity)

      if got != tt.Output {
        t.Errorf("got %d, want %d", got, tt.Output)
      }
    }
  })
}

func TestFindKnapsackedItems(t *testing.T) {
  t.Skip()
}
