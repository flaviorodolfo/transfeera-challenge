package domain

import "context"

type RecebedorRepository interface {
	CriarRecebedor(ctx context.Context, recebedor *Recebedor) error
}
