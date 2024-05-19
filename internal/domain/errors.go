package domain

import "errors"

var (
	ErrEmailInvalido             = errors.New("email inválido")
	ErrNomeInvalido              = errors.New("Nome inválido")
	ErrChaveInvalida             = errors.New("chave inválida")
	ErrTipoChaveInvalida         = errors.New("tipo chave inválida")
	ErrCpfInvalido               = errors.New("cpf inválido")
	ErrCnpjInvalido              = errors.New("cnpj inválido")
	ErrRecebedorNaoEncontrado    = errors.New("recebedor não existe")
	ErrRecebedorNaoPermiteEdicao = errors.New("recebedor com status Validado apenas permite edicao de email")
)
