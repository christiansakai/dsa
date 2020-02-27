package solution

func Solve(num1, num2 string) string {
	result := []int{}

	for i := len(num1) - 1; i >= 0; i-- {
		row := []int{}
		carry := 0

		for j := len(num2) - 1; j >= 0; j-- {
			product := (num1[i] * num2[j]) + carry

			carry = product / 10
			remainder := product % 10

			row = apend(row, remainder)

			if carry > 0 {
				row = append(row, carry)
			}
		}

		result = append(result, row)
	}

}
