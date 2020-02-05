package solution

func TopDown(set []int, S int) int {
  if len(set) == 0 || S == 0 {
    return 0
  }

  cache := map[int]map[int]int{}

  return recurse(set, S, len(set) - 1, cache)
}

func recurse(set []int, S, index int, cache map[int]map[int]int) int {
  if index == - 1 {
    if S == 0 {
      return 1
    }

    return 0
  }

  if _, ok := cache[S]; ok {
    if result, ok := cache[S][index]; ok {
      return result
    }
  }

  plus := recurse(set, S - set[index], index - 1, cache)
  minus := recurse(set, S + set[index], index - 1, cache)

  result := plus + minus

  if _, ok := cache[S]; !ok {
    cache[S] = map[int]int{}
  }

  cache[S][index] = result

  return result
}
