package solution

import "math"

func TopDown(cuttings []int, length int) int {
	if len(cuttings) == 0 || length ==0 {
		return 0
	}

	return recurse(cuttings, length, len(cuttings) - 1)
}

func recurse(cuttings []int, length, index int) int {
	if length == 0 || index == -1 {
		return 0
	}

	var max float64 = 0

  if cuttings[index] <= length {
    with := 1 + recurse(cuttings, length - cuttings[index], index)
    max = math.Max(max, float64(with))
  }

  without := recurse(cuttings, length, index - 1)
  max = math.Max(max, float64(without))

	result := int(max)

	return result
}
