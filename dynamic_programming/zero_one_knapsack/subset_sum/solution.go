package solution

func TopDown(set []int, target int) bool {
  if len(set) == 0 {
    return false
  }

  cache := map[int]map[int]bool{}

  return recurse(set, len(set) - 1, target, cache)
}

func recurse(set []int, index, target int, cache map[int]map[int]bool) bool {
  if index == 0 {
    return set[0] == target
  }

  if _, ok := cache[index]; ok {
    if result, ok := cache[index][target]; ok {
      return result
    }
  }

  with := recurse(set, index - 1, target - set[index], cache)
  without := recurse(set, index - 1, target, cache)

  result := with || without

  if _, ok := cache[index]; !ok {
    cache[index] = map[int]bool{}
  }

  cache[index][target] = result

  return result
}

func BottomUp(set []int, target int) bool {
  if len(set) == 0 {
    return false
  }

  equal := make([][]bool, len(set))

  // With 0 target we can't fit any subset that fits
  for i := 0; i < len(equal); i++ {
    equal[i] = make([]bool, target + 1) 
    equal[i][0] = false
  }

  // With only 1 item in set, we also can't make two subsets
  for t := 0; t <= target; t++ {
    equal[0][t] = false
  }

  // Process the rest of the set for the rest of the target
  for i := 1; i < len(set); i++ {
    for t := 1; t <= target; t++ {
      with := equal[i - 1][t - i]
      without := equal[i - 1][t]

      equal[i][t] = with || without
    }
  }

  return equal[len(set) - 1][target]
}

func BottomUpMinSpace(set []int, target int) bool {
  if len(set) == 0 {
    return false
  }

  equal := make([]bool, target + 1)

  // With 0 target we can't fit any subset that fits
  equal[0] = false

  // Process the rest of the set for the rest of the target
  for i := 0; i < len(set); i++ {
    for t := 1; t <= target; t++ {
      with := true || equal[t - i]
      without := equal[t]

      equal[t] = with || without
    }
  }

  return equal[target]
}
