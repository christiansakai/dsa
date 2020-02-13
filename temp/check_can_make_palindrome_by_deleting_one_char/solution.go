package solution

import "math"

func SolveTopDown(s string) bool {
	if len(s) <= 2 {
		return true
	}

	dict := map[byte]int{}

}

// func SolveTopDown(s string) bool {
//   if len(s) <= 2 {
//     return true
//   }

//   cache := map[int]map[int]int{}
//   lps := solveLPSRecursive(s, 0, len(s) - 1, cache)

//   remainder := len(s) - lps

//   return remainder <= 1
// }

// func solveLPSRecursive(s string, start, end int, cache map[int]map[int]int) int {
//   if start > end {
//     return 0
//   }

//   if start == end {
//     return 1
//   }

//   if _, ok := cache[start]; ok {
//     if result, ok := cache[start][end]; ok {
//       return result
//     }
//   }

//   if s[start] == s[end] {
//     return 2 + solveLPSRecursive(s, start + 1, end - 1, cache)
//   }

//   skipStart := solveLPSRecursive(s, start + 1, end, cache)
//   skipEnd := solveLPSRecursive(s, start, end - 1, cache)

//   max := int(math.Max(
//     float64(skipStart),
//     float64(skipEnd),
//   ))

//   if _, ok := cache[start]; !ok {
//     cache[start] = map[int]int{}
//   }

//   cache[start][end] = max

//   return max
// }
