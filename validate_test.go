package cpf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	var check_cpf bool

  check_cpf = Validate("111.444.777-35")
	assert.Equal(t, true, check_cpf, "should be 'true'")

	check_cpf = Validate("218.425.985-38")
	assert.Equal(t, true, check_cpf, "should be 'true'")

	check_cpf = Validate("218.425.985-42")
	assert.Equal(t, false, check_cpf, "should be 'false'")
}
