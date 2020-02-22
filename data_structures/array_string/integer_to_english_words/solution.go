package solution

func Solve(num int) string {
	result := []string{}

	numstr := intToString(num)
	for i := 0; i < len(numstr); i++ {
		var text string

		if i+1 < len(numstr) {
			text = getStringForTwoDigits(numstr[i], numstr[i+1])
		} else {
			text = getStringForOneDigit(numstr[i])
		}

		tensString := getTensString(i)
		if tensString != "" {
			text += " " + tensString
		}

		numstr = append(numstr, text)
	}

	return result
}

var oneDigitSpecial map[byte]string = map[byte]string{
	'0': "Zero",
	'1': "One",
	'2': "Two",
	'3': "Three",
	'4': "Four",
	'5': "Five",
	'6': "Six",
	'7': "Seven",
	'8': "Eight",
	'9': "Nine",
}

func getStringForTwoDigits(d1, d2 byte) string {
	str := string([]byte{d1, d2})

	if result, ok := twoDigitSpecial[str]; ok {
		return result
	}

}

var twoDigitSpecial map[string]string = map[string]string{
	"10": "Ten",
	"11": "Eleven",
	"12": "Twelve",
	"13": "Thirteen",
	"15": "Fifteen",
	"20": "Twenty",
	"30": "Thirty",
	"40": "Forty",
	"50": "Fifty",
}
