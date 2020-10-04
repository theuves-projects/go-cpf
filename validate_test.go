package cpf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	var checkCpf bool

  checkCpf = Validate("111.444.777-35")
	assert.Equal(t, true, checkCpf, "should be 'true'")

	checkCpf = Validate("218.425.985-38")
	assert.Equal(t, true, checkCpf, "should be 'true'")

	checkCpf = Validate("218.425.985-42")
	assert.Equal(t, false, checkCpf, "should be 'false'")
}

func TestParseCpf(t *testing.T) {
	cpfNumber, _ := parseCpf("111.444.777-35")
	expectedDigits := [9]int{1, 1, 1, 4, 4, 4, 7, 7, 7}
	expectedCheckDigits := [2]int{3, 5}

	assert.Equal(t, expectedDigits, cpfNumber.digits, "should be CPF digits")
	assert.Equal(t, expectedCheckDigits, cpfNumber.check_digits, "should be CPF check digits")
}

func TestGetAllDigits(t *testing.T) {
	param := getAllDigits("111.444.777-35")
	expected := []int{1, 1, 1, 4, 4, 4, 7, 7, 7, 3, 5}

	assert.Equal(t, expected, param, "should get all digits")
}
