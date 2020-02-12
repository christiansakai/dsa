package solution

import "math"

// import "fmt"

func Solve(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	b := []byte(s)
	set := map[byte]bool{}
	// foundRepeat := false

	var max float64 = 0
	for i := 0; i < len(b); i++ {
		for j := i; j < len(b); j++ {
			if _, ok := set[b[j]]; !ok {
				// fmt.Println("HERE", j)

				set[b[j]] = true
				continue
			}

			// fmt.Println("THEREHERE", j)
			// foundRepeat = true
			length := j - i
			max = math.Max(max, float64(length))
		}
	}

	// if foundRepeat {
	return int(max)
	// }

	// return len(s)
}
