func lengthOfLongestSubstring(s string) int {
  if len(s) <= 1 {
    return len(s)
  }
  
  b := []byte(s)
  set := map[byte]bool{}
  
  var max float64 = 0
  for i := 0; i < len(b); i++ {
    for j := i; j < len(b); j++ {
      if _, ok := set[b[j]]; !ok {
        set[b[j]] = true
        continue
      }
      
      length := j - i
      max = math.Max(max, float64(length))
      i = j
    }
  }

  return int(max)
}
