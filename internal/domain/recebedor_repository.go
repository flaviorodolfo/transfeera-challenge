package domain

type RecebedorRepository interface {
	CriarRecebedor(recebedor *Recebedor) error
}
