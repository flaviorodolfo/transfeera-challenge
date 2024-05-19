package domain

type RecebedorRepository interface {
	BuscarRecebedorPorID(id uint) (*Recebedor, error)
	CriarRecebedor(recebedor *Recebedor) error
	EditarRecebedor(recebedor *Recebedor) error
}
