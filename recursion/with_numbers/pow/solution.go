package solution

import "math"

func Solve(x float64, n int) float64 {
	cache := map[int]float64{}
	nAbs := int(math.Abs(float64(n)))
	result := recurse(x, nAbs, cache)

	if n < 0 {
		return 1 / result
	}

	return result
}

func recurse(x float64, n int, cache map[int]float64) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if result, ok := cache[n]; ok {
		return result
	}

	var result float64
	if n%2 == 0 {
		subProb := recurse(x, n/2, cache)
		result = subProb * subProb
	} else {
		subProb := recurse(x, n-1, cache)
		result = x * subProb
	}

	cache[n] = result

	return result
}
