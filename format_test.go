package cpf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	var expected string = "111.444.777-35"
  var cpf string

  cpf, _ = Format("11144477735")
	assert.Equal(t, expected, cpf, "should format CPF")

  cpf, _ = Format("111.444.77735")
	assert.Equal(t, expected, cpf, "should format CPF")

  cpf, _ = Format("111444777-35")
	assert.Equal(t, expected, cpf, "should format CPF")
}
