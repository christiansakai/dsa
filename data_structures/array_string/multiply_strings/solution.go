package solution

func Solve(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]byte, len(num1)+len(num2))
	for i := 0; i < len(result); i++ {
		res[i] = '0'
	}

	for i := len(num2) - 1; i >= 0; i-- {
	}

}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := []rune(num1 + num2)
	for i := 0; i < len(num1)+len(num2); i++ {
		res[i] = '0'
	} // the above code prepares res as a string of "000...0" of length len(num1)+ len(num2).

	for i := len(num2) - 1; i >= 0; i-- {
		carry := 0
		j := len(num1) - 1
		m := 0
		for carry != 0 || j >= 0 {
			m = carry + int(res[i+j+1]-'0')
			if j >= 0 {
				m += int(num2[i]-'0') * int(num1[j]-'0')
			}
			res[i+j+1] = rune(m%10 + '0')
			carry = m / 10
			j--
		}
	}
	count := 0 // to count how many zeros are to be removed in the front of "res"
	for i := 0; i < len(res) && res[i] == '0'; i++ {
		count++
	}
	return string(res[count:])
}
