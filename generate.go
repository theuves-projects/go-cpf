package cpf

import (
	"math/rand"
	"strconv"
	"time"
)

// Generate a random CPF number
func Generate() string {
	var (
		values      [9]int
		checkDigits [2]int
		digits      string
	)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 9; i++ {
		values[i] = rand.Intn(10)
		digits += strconv.Itoa(values[i])
	}

	checkDigits = getCheckDigits(values)

	digits += strconv.Itoa(checkDigits[0])
	digits += strconv.Itoa(checkDigits[1])

	cpf, _ := Format(digits)
	return cpf
}
