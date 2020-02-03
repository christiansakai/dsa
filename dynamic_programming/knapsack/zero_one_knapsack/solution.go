package solution

import "math"

type Fruit struct {
  Name string
  Weight int
  Profit int
}

func TopDown(fruits []Fruit, capacity int) int {
  if len(fruits) == 0 {
    return 0
  }

  cache := map[int]map[int]int{}

  return recurse(fruits, len(fruits) - 1, capacity, cache)
}

func recurse(fruits []Fruit, index, capacity int, cache map[int]map[int]int) int {
  if _, ok := cache[index]; ok {
    if result, ok := cache[index][capacity]; ok {
      return result
    }
  }

  if index == 0 {
    if fruits[0].Weight <= capacity {
      return fruits[0].Profit
    }

    return 0
  }

  var max float64 = 0

  if fruits[index].Weight <= capacity {
    with := fruits[index].Profit + recurse(
      fruits,
      index - 1,
      capacity - fruits[index].Weight,
      cache,
    )

    max = math.Max(max, float64(with))
  }

  without := recurse(fruits, index - 1, capacity, cache)
  max = math.Max(max, float64(without))

  result := int(max)

  if _, ok := cache[index]; !ok {
    cache[index] = map[int]int{}
  }

  cache[index][capacity] = result

  return result
}

func BottomUp(fruits []Fruit, capacity int) int {
  if len(fruits) == 0 {
    return 0
  }

  profits := make([][]int, len(fruits))

  // With 0 capacity we have 0 profit
  for i := 0; i < len(fruits); i++ {
    profits[i] = make([]int, capacity + 1)
    profits[i][0] = 0
  }

  // With only 1 fruit, take it if it is not more than capacity
  for c := 0; c <= capacity; c++ {
    if fruits[0].Weight <= c {
      profits[0][c] = fruits[0].Profit
    } else {
      profits[0][c] = 0
    }
  }

  // Process the rest of sub-arrays for the rest of the capacities
  for i := 1; i < len(fruits); i++ {
    for c := 1; c <= capacity; c++ {
      var max float64 = 0

      if fruits[i].Weight <= c {
        with := fruits[i].Profit + profits[i - 1][c - fruits[i].Weight]
        max = math.Max(max, float64(with))
      }

      without := profits[i - 1][c]
      max = math.Max(max, float64(without))

      profits[i][c] = int(max)
    }
  }

  return profits[len(fruits) - 1][capacity]
}

func BottomUpMinSpace(fruits []Fruit, capacity int) int {
  if len(fruits) == 0 {
    return 0
  }

  profits := make([]int, capacity + 1)

  for i := 0; i < len(fruits); i++ {
    // Run the iteration in reverse
    for c := capacity; c >= 0; c-- {
      var max float64 = 0

      if fruits[i].Weight <= c {
        with := fruits[i].Profit + profits[c - fruits[i].Weight]
        max = math.Max(max, float64(with))
      }

      without := profits[c]
      max = math.Max(max, float64(without))

      profits[c] = int(max)
    }
  }

  return profits[capacity]
}

func FindKnapsackedItems(fruits []Fruit, capacity int) []string {
  return []string{}
}
