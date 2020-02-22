package solution

import "math"

func Solve(str string) int {
	if len(str) == 0 {
		return 0
	}

	i := 0
	isNegative := false

	// Skip whitespace
	for ; i < len(str) && str[i] == ' '; i++ {
	}
	if i == len(str) {
		return 0
	}

	// Check negativity
	if str[i] == '-' {
		isNegative = true
		i++
	} else if str[i] == '+' {
		i++
	}

	if i == len(str) {
		return 0
	}

	// Collect integer
	integerArr := []byte{}
	for ; i < len(str) && isInteger(str[i]); i++ {
		integerArr = append(integerArr, str[i])
	}

	return convertToInteger(integerArr, isNegative)
}

func isInteger(b byte) bool {
	return b >= 48 && b <= (48+9)
}

func convertToInteger(intArr []byte, isNegative bool) int {
	integer := 0

	for i := 0; i < len(intArr); i++ {
		d := bToD(intArr[i])
		integer *= 10
		integer += d

		if integer > math.MaxInt32 {
			if isNegative {
				return math.MinInt32
			}

			return math.MaxInt32
		}
	}

	if isNegative {
		return -1 * integer
	}

	return integer
}

func bToD(b byte) int {
	return int(b - 48)
}
