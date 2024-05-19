package domain

type RecebedorRepository interface {
	BuscarRecebedorPorId(id uint) (*Recebedor, error)
	CriarRecebedor(recebedor *Recebedor) error
	EditarRecebedor(recebedor *Recebedor) error
	EditarEmailRecebedor(id uint, email string) error
}
