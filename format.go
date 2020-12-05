package cpf

import (
	"errors"
	"regexp"
)

// Format a CPF number (mask 999.999.999-99).
func Format(cpf string) (string, error) {
	r := regexp.MustCompile(`^(\d{1,3})\.?(\d{1,3})\.?(\d{1,3})-?(\d{1,2})$`)

	if !r.MatchString(cpf) {
		return "", errors.New("invalid CPF number")
	}

	return r.ReplaceAllString(cpf, "$1.$2.$3-$4"), nil
}
