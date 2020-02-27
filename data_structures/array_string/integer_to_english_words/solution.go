package solution

func Solve(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := []string{}

	divider := 10000000000
	for num > 0 {
		n := num / divider
		if n >= 1 && n <= 9 {
			intStr := lessThanTen[n]
			result = append(result, intStr)

			zeroes := hundredAbove[divider]
			result = append(result, zeroes)
		} else {

		}

		num = num % divider

		divider /= 1000
	}

	return result
}

var lessThanTen map[int]string = map[int]string{
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

var lessThanTwenty map[string]string = map[string]string{
	"11": "Eleven",
	"12": "Twelve",
	"13": "Thirteen",
	"14": "Fourteen",
	"15": "Fifteen",
	"16": "Sixteen",
	"17": "Seventeen",
	"18": "Eighteen",
	"19": "Nineteen",
}

var lessThanHundred map[string]string = map[string]string{
	"10": "Ten",
	"20": "Twenty",
	"30": "Thirty",
	"40": "Forty",
	"50": "Fifty",
	"60": "Sixty",
	"70": "Seventy",
	"80": "Eighty",
	"90": "Ninety",
}

var hundredAbove map[int]string = map[int]string{
	100:        "Hundred",
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}
