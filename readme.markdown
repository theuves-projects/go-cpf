# go-cpf ![GoDoc](https://godoc.org/github.com/theuves/go-cpf?status.svg)

> Funções básicas para a manipulação de números de [CPF](https://pt.wikipedia.org/wiki/Cadastro_de_pessoas_f%C3%ADsicas).

## Instalação

Com `go-get`:

```bash
$ go get github.com/theuves/go-cpf
```

Veja a [documentação aqui](https://godoc.org/github.com/theuves/go-cpf) (ou use `$ godoc github.com/theuves/go-cpf`).

## Exemplos

Veja alguns exemplos abaixo:

```go
package main

import (
    "fmt"
    "github.com/theuves/go-cpf"
)

func main() {
    fmt.Println(cpf.Format("01234567890")) // 012.345.678-90, <nil>
    fmt.Println(cpf.Generate()) // retorna um número de CPF aleaatório válido
    fmt.Println(cpf.Validate("012.345.678-90")) // true
}
```

## Build status

| Branch | Badge |
|:---:|:---:|
| `master` | ![Go](https://github.com/theuves/go-cpf/workflows/Go/badge.svg?branch=master) |
| `dev` | ![Go](https://github.com/theuves/go-cpf/workflows/Go/badge.svg?branch=dev) |

## Licença

MIT &copy; [Matheus Alves](https://github.com/theuves)
