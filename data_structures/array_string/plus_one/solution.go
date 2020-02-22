package solution

func Solve(digits []int) []int {
	carry := 0
	newDigits := make([]int, len(digits))

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]

		var sum int
		if i == len(digits)-1 {
			sum = digit + 1 + carry
		} else {
			sum = digit + carry
		}

		carry = 0

		if sum >= 10 {
			carry = 1
			remainder := sum - 10

			newDigits[i] = remainder
		} else {
			newDigits[i] = sum
		}
	}

	if carry > 0 {
		newDigits = append([]int{carry}, newDigits...)
	}

	return newDigits
}
