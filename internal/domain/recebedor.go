package domain

type TipoChavePix string

const (
	Cpf            TipoChavePix = "CPF"
	Cnpj           TipoChavePix = "CNPJ"
	Email          TipoChavePix = "EMAIL"
	Telefone       TipoChavePix = "TELEFONE"
	ChaveAleatoria TipoChavePix = "CHAVE_ALEATORIA"
)

type Recebedor struct {
	Id           uint         `json:"id" `
	CpfCnpj      string       `json:"cpf_cnpj" validate:"required" `
	Nome         string       `json:"nome" validate:"required"`
	TipoChavePix TipoChavePix `json:"tipo_chave_pix" validate:"required"`
	ChavePix     string       `json:"chave_pix" validate:"required"`
	Status       string       `json:"status"`
	Email        string       `json:"email"`
}
