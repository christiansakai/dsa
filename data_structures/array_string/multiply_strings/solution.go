package solution

import "fmt"

func Solve(num1, num2 string) string {
	bs := [][]byte{}

	for i := len(num1) - 1; i >= 0; i-- {
		d1 := num1[i]
		multipliedWithSingleDigit := multiplyNumWithSingleDigit(
			num2,
			d1,
		)

		numOfZeroes := len(num1) - 1 - i
		for j := 0; j < numOfZeroes; j++ {
			multipliedWithSingleDigit = append(
				[]byte{'0'},
				multipliedWithSingleDigit...,
			)
		}

		bs = append(bs, multipliedWithSingleDigit)
	}

	sum := sumEverything(bs)

	fmt.Println(string(sum))
	reversed := reverse(sum)

	return string(reversed)
}

func multiplyNumWithSingleDigit(num string, d byte) []byte {
	b := []byte{}
	carry := 0

	for i := len(num) - 1; i >= 0; i-- {
		d1 := byteToInt[byte(num[i])]
		d2 := byteToInt[byte(d)]

		multiplied := (d1 * d2) + carry
		if multiplied >= 10 {
			carry = multiplied / 10
		}

		remainder := multiplied % 10

		b = append(b, intToByte[remainder])
	}

	if carry > 0 {
		b = append(b, intToByte[carry])
	}

	return b
}

var byteToInt map[byte]int = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

var intToByte map[int]byte = map[int]byte{
	0: '0',
	1: '1',
	2: '2',
	3: '3',
	4: '4',
	5: '5',
	6: '6',
	7: '7',
	8: '8',
	9: '9',
}

func reverse(b []byte) []byte {
	mid := len(b) / 2
	for i := 0; i < mid; i++ {
		temp := b[i]
		b[i] = b[len(b)-1-i]
		b[len(b)-1-i] = temp
	}

	return b
}

func sumEverything(bs [][]byte) []byte {
	b := []byte{}
	carry := 0

	for _, row := range bs {
		sum := 0 + carry
		for _, dbyte := range row {
			sum += byteToInt[dbyte]
		}

		if sum >= 10 {
			carry = sum / 10
		}

		remainder := sum % 10
		b = append(b, intToByte[remainder])
	}

	return b
}
