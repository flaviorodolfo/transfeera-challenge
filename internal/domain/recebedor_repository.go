package domain

type RecebedorRepository interface {
	BuscarRecebedorPorId(id uint) (*Recebedor, error)
	BuscarRecebedoresPorCampo(valor, nomeCampo string, offset int) ([]*Recebedor, error)
	ContarRecebedoresPorCampo(valor, nomeCampo string) (int, error)
	CriarRecebedor(recebedor *Recebedor) error
	EditarRecebedor(recebedor *Recebedor) error
	EditarEmailRecebedor(id uint, email string) error
}
