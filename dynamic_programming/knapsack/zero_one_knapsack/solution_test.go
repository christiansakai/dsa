package solution

import (
  "testing"
  "reflect"
)

func TestSolution(t *testing.T) {
  tests := []struct{
    Input struct{
      Fruits []Fruit
      Capacity int
    }

    Output int
  }{
    {
      Input: struct{
        Fruits []Fruit
        Capacity int
      }{
        Fruits: []Fruit{
          Fruit{"Apple", 2, 4},
          Fruit{"Orange", 3, 5},
          Fruit{"Banana", 1, 3},
          Fruit{"Melon", 4, 7},
        },
        Capacity: 5,
      },
      Output: 10,
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

  t.Run("Bottom-up Dynamic Programming with Minimum Space", func(t *testing.T) {
    for _, tt := range tests {
    got := BottomUpMinSpace(tt.Input.Fruits, tt.Input.Capacity)
      if got != tt.Output {
        t.Errorf("got %d, want %d", got, tt.Output)
      }
    }
  })
}

func TestFindKnapsackedItem(t *testing.T) {
  t.Skip()

  fruits := []Fruit{
    Fruit{"Apple", 2, 4},
    Fruit{"Orange", 3, 5},
    Fruit{"Banana", 1, 3},
    Fruit{"Melon", 4, 7},
  }

  capacity := 5

  got := FindKnapsackedItems(fruits, capacity)
  want := []string{"Apple"}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v, want %v", got, want)
  }
}
