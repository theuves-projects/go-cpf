package cpf

import (
    "errors"
    "math/rand"
    "regexp"
    "strconv"
    "strings"
    "time"
)

// Calcular os digitos verificadores
// de um número de CPF.
func digits(nums string) string {
    var cpf = []int{}
    scpf := strings.Split(nums, "")

    var v1, v2 int

    for i := 0; i < len(scpf); i++ {
        if n, err := strconv.Atoi(scpf[i]); err == nil {
            cpf = append(cpf, n)
        }
    }

    for i := 0; i < 9; i++ {
        v2 = v2 + cpf[i] * (9 - (i % 10))
        v1 = v1 + cpf[i] * (9 - ((i + 1) % 10))
    }

    v2 = (v2 % 11) % 10
    v1 = ((v1 + v2 * 9) % 11) % 10

    return strconv.Itoa(v1) + strconv.Itoa(v2)
}

// Formatar qualquer número de CPF.
func Format(cpf string) (string, error) {
    r := regexp.MustCompile(`^(\d{1,3})\.?(\d{1,3})\.?(\d{1,3})-?(\d{1,2})$`)

    if !r.MatchString(cpf) {
        return "", errors.New("cpf.Format: CPF invalid")
    } else {
        return r.ReplaceAllString(cpf, "$1.$2.$3-$4"), nil
    }
}

// Gerar um número de CPF aleatório válido.
func Generate() string {
    rand.Seed(time.Now().UTC().UnixNano())

    n := rand.Perm(10)[:9]
    var s []string

    for i := 0; i < 9; i++ {
        s = append(s, strconv.Itoa(n[i]))
    }

    s = append(s, digits(strings.Join(s, "")))

    var ret string

    if cpf, err := Format(strings.Join(s, "")); err == nil {
        ret = cpf
    }

    return ret
}

// Válidar um número de CPF.
func Validate(cpf string) bool {
    str, err := Format(cpf);

    if (err != nil) {
        return false
    } else {
        // Para verificar se todos os dígitos são iguais.
        r1 := regexp.MustCompile(`^(1+|2+|3+|4+|5+|6+|7+|8+|9+)$`)

        // Para desformatar.
        r2 := regexp.MustCompile(`\D+`)

        // Verifica se todos os dígitos do CPF são iguais
        // caso seja retorna `false`.
        if r1.MatchString(r2.ReplaceAllString(str, "")) {
            return false
        }

        s := strings.Split(str, "-")

        // Verifica se os dígitos verificadores
        // informados são válidos.
        if s[1] == digits(s[0]) {
            return true
        } else {
            return false
        }
    }
}