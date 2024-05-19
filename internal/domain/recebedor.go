package domain

type TipoChavePix string

const (
	Cpf            TipoChavePix = "CPF"
	Cnpj           TipoChavePix = "CNPJ"
	Email          TipoChavePix = "EMAIL"
	Telefone       TipoChavePix = "TELEFONE"
	ChaveAleatoria TipoChavePix = "CHAVE_ALEATORIA"
)

type PaginaRecebedores struct {
	Total        int          `json:"total"`
	PorPagina    int          `json:"por_pagina"`
	PaginaAtual  int          `json:"pagina_atual"`
	TotalPaginas int          `json:"total_paginas"`
	Recebedores  []*Recebedor `json:"recebedores"`
}

type Recebedor struct {
	Id           uint         `json:"id" `
	CpfCnpj      string       `json:"cpf_cnpj" validate:"required" `
	Nome         string       `json:"nome" validate:"required"`
	TipoChavePix TipoChavePix `json:"tipo_chave_pix" validate:"required"`
	ChavePix     string       `json:"chave_pix" validate:"required"`
	Status       string       `json:"status"`
	Email        string       `json:"email"`
}
