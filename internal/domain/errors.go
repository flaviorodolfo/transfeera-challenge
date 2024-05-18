package domain

import "errors"

var (
	ErrEmailInvalido = errors.New("email inválido")
	ErrChaveInvalida = errors.New("chave inválida")
	ErrCpfInvalido   = errors.New("cpf inválido")
	ErrCnpjInvalido  = errors.New("cnpj inválido")
)
