package solution

// import "math"
// import "fmt"

func TopDown(str string) int {
  if len(str) == 0 {
    return 0
  }

  byteStr := []byte(str)

  return recurse(byteStr, 0, len(byteStr) - 1)
}

func recurse(str []byte, start, end int) int {
  if start > end {
    return 0
  }

  if start == end {
    return 1
  }

  if str[start] == str[end] {
    subPalindromesCount := recurse(str, start + 1, end - 1) 

    if subPalindromesCount > 0 {
      return subPalindromesCount + 1
    }

    return 0
  }

  total := 0

  skipFront := recurse(str, start + 1, end)
  total += skipFront

  skipBack := recurse(str, start, end - 1)
  total += skipBack

  return total
}
