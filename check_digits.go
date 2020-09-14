package cpf

func reverseDigits(digits [9]int) [9]int {
	var invertedDigits [9]int
	const maxLength = 8

	for i := maxLength; i >= 0; i-- {
		invertedDigits[maxLength - i] = digits[i]
	}

	return invertedDigits
}

func getCheckDigits(digits [9]int) [2]int {
	var (
		v1 int
		v2 int
	)

	digits = reverseDigits(digits)

	for i := 0; i < 9; i++ {
		v1 = v1 + digits[i] * (9 - (i % 10))
		v2 = v2 + digits[i] * (9 - ((i + 1) % 10))
	}

	v1 = (v1 % 11) % 10
	v2 = ((v2 + v1*9) % 11) % 10

	return [2]int{v1, v2}
}
