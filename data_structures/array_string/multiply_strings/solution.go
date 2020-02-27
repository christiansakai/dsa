package solution

func Solve(num1, num2 string) string {
  result := []int{}
  carry := 0

  for i := len(num1) - 1; i >= 0; i-- {
    for j := len(num2) - 1; j >= 0; j-- {
      num1Digit := num1[i]
      num2Digit := num2[j]

      product := (byteToInt(num1Digit) * byteToInt(num2Digit)) + carry
      productIndex := i + j

      if !(productIndex < len(result)) {
      }
    }
  }
}

class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        result, carry = [], 0
        for num1Idx, num1Digit in enumerate(num1[::-1]):
            for num2Idx, num2Digit in enumerate(num2[::-1]):
                digitProduct = (ord(num1Digit) - ord('0')) * (ord(num2Digit) - ord('0')) + carry
                resultIdx = num1Idx + num2Idx
                if not resultIdx < len(result):
                    result.append(0)
                # imagine if you were summing the results
                # as you did the multiplication...
                result[resultIdx] += digitProduct
                carry = result[resultIdx] // 10
                result[resultIdx] %= 10
            if carry:
                result.append(carry)
                carry = 0
        resultStr = "".join(map(str,result[::-1])).lstrip("0")
        return resultStr if resultStr else "0"
