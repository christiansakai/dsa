package solution

import "strings"

func Solve(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := []string{}

	divider := 1000000000
	for num > 0 {
		n := num / divider
		if n > 0 {
			solveForThreeDigit(n, &result)

			zeroes := thousandAndAbove[divider]
			if zeroes != "" {
				result = append(result, zeroes)
			}
		}

		num = num % divider
		divider /= 1000
	}

	return strings.Join(result, " ")
}

func solveForThreeDigit(num int, result *[]string) {
	if num <= 10 {
		*result = append(*result, oneToTen[num])
	} else if num < 20 {
		*result = append(*result, elevenToNineteen[num])
	} else if num < 100 {
		tens := num / 10
		*result = append(*result, twentyToNinety[tens])

		ones := num % 10
		if ones > 0 {
			*result = append(*result, oneToTen[ones])
		}
	} else {
		hundreds := num / 100
		*result = append(*result, oneToTen[hundreds])
		*result = append(*result, "Hundred")

		tens := num % 100
		if tens > 0 {
			solveForThreeDigit(tens, result)
		}
	}
}

var oneToTen map[int]string = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

var elevenToNineteen map[int]string = map[int]string{
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}

var twentyToNinety map[int]string = map[int]string{
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

var thousandAndAbove map[int]string = map[int]string{
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}
