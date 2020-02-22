package solution

func Solve(nums []int) int {
  dict := map[int]bool{}

  for _, n := range nums {
    if dict[n] == false {
      dict[n] = true
    } else {
      dict[n] = false
    }
  }

  for k, v := range dict {
    if v {
      return k
    }
  }

  return -1
}
