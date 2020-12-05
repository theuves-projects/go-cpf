package cpf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	var randomCpf string = Generate()

	assert.Equal(t, true, Validate(randomCpf), "should generate a valid CPF")
}
