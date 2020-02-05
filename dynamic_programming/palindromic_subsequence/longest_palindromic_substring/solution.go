package solution

import "math"
import "fmt"

func TopDown(str string) int {
	if len(str) == 0 {
		return 0
	}

	byteStr := []byte(str)

  return recurse2(byteStr, 0, 0, "")
}

func recurse2(str []byte, i, j int, acc string) int {
  if i == len(str) {
    return 0
  }

  if j == len(str) {
    fmt.Println(str)
    var max float64 = 0
    if isPalindrome(acc) {
      max = float64(len(acc))
    }

    nextI := i + 1
    nextJ := nextI + 1

    iPlus := recurse2(str, nextI, nextJ, "") 
    max = math.Max(max, float64(iPlus))

    return int(max)
  }

  jPlus := recurse2(str, i, j + 1, acc + string([]byte{str[j]}))
  // iPlus := recurse2(str, i + 1, j, acc)

  // result := math.Max(float64(jPlus), float64(iPlus))

  return int(jPlus)
}

// func recurse(str []byte, index int, acc string) int {
//   if index == len(str) {
//     fmt.Println(acc)
//     if isPalindrome(acc) {
//       return len(acc)
//     }

//     return 0
//   }

//   without := recurse(str, index + 1, acc)
//   with := recurse(str, index + 1, acc + string([]byte{str[index]}))

//   result := math.Max(float64(with), float64(without))

//   return int(result)
// }

func isPalindrome(str string) bool {
  b := []byte(str)
  length := len(b) / 2

  for i := 0; i < length; i++ {
    if b[i] != b[len(b) - 1 - i] {
      return false
    }
  }

  return true
}
