package cpf

import (
  "errors"
  "regexp"
  "strconv"
)

type CpfNumber struct {
  digits [9]int
  check_digits [2]int
}

// Validate a CPF number
func Validate(cpf string) bool {
  cpfParsed, err := parseCpf(cpf)

  if err != nil {
    return false
  }

  var check_digits [2]int = getCheckDigits(cpfParsed.digits)

  return check_digits == cpfParsed.check_digits
}

// Parse a CPF number
func parseCpf(cpf string) (CpfNumber, error) {
  var allDigits []int = getAllDigits(cpf)
  var cpfNumber CpfNumber = CpfNumber{}

  if len(allDigits) != 11 {
    return cpfNumber, errors.New("invalid CPF number")
  }

  for i := 0; i < cap(cpfNumber.digits); i++ {
    cpfNumber.digits[i] = allDigits[:9][i]
  }
  for i := 0; i < cap(cpfNumber.check_digits); i++ {
    cpfNumber.check_digits[i] = allDigits[9:][i]
  }

  return cpfNumber, nil;
}

// Get only digits ([]int) from a string
func getAllDigits(str string) []int {
  r := regexp.MustCompile(`\D+`)
  return stringsToInts(r.ReplaceAllString(str, ""))
}

// Convert string of digits to []int
func stringsToInts(str string) []int {
	var arrayInt = []int{}
	var value int

	for i := 0; i < len(str); i++ {
		value, _ = strconv.Atoi(string(str[i]))
		arrayInt = append(arrayInt, value)
	}

	return arrayInt
}
