package cpf

import (
  "errors"
  "regexp"
  "strconv"
)

// Validate a CPF number
func Validate(cpf string) bool {
  var allDigits []int = getAllDigits(cpf)

  if (len(allDigits) != 11) {
    return false
  }

  var digits []int = allDigits[:9]
  var check_digits []int = getCheckDigits(digits)

  return check_digits[0] == allDigits[10] && check_digits[1] == allDigits[11]
}

// Get only digits ([]int) from a string
func getAllDigits(str string) string {
  r := regexp.MustCompile(`\D+`)
  return stringsToInts(r.ReplaceAllString(str, ""))
}

// Convert string to []int
func stringsToInts(str string) []int {
	var arrayInt = []int{}
	var value int

	for i := 0; i < len(str); i++ {
		value, _ = strconv.Atoi(string(str[i]))
		arrayInt = append(arrayInt, value)
	}

	return arrayInt
}
