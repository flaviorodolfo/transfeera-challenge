package domain

import (
	"errors"
	"fmt"
)

type ErrRecebedoresNaoDeletados struct {
	IdsSemSucesso []uint `json:"ids_sem_sucesso"`
	IdsComSucesso []uint `json:"ids_com_sucesso"`
}

func (e ErrRecebedoresNaoDeletados) Error() string {
	return fmt.Sprintf("deletados: %v não deletados:%v", e.IdsComSucesso, e.IdsSemSucesso)
}

var (
	ErrEmailInvalido             = errors.New("email inválido")
	ErrNomeInvalido              = errors.New("nome inválido")
	ErrChaveInvalida             = errors.New("formato de chave inexistente")
	ErrTipoChaveInvalida         = errors.New("tipo de chave inválida")
	ErrChaveTipoNaoCorresponde   = errors.New("chave não corresponde com o tipo")
	ErrCpfInvalido               = errors.New("cpf inválido")
	ErrCnpjInvalido              = errors.New("cnpj inválido")
	ErrRecebedorNaoEncontrado    = errors.New("recebedor não existe")
	ErrRecebedorNaoPermiteEdicao = errors.New("recebedor com status Validado apenas permite edicao de email")
)
