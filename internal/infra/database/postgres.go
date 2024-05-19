package database

import (
	"database/sql"

	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
)

type postgresRecebedorRepository struct {
	DB *sql.DB
}

func NewPostgresRecebedorRepository(db *sql.DB) *postgresRecebedorRepository {
	return &postgresRecebedorRepository{DB: db}
}

func (r *postgresRecebedorRepository) CriarRecebedor(recebedor *domain.Recebedor) error {
	query := "INSERT INTO pagamento.recebedores (cpf_cnpj, nome, tipo_chave_pix,chave_pix, status_recebedor, email) VALUES ($1, $2, $3,$4, $5,$6) RETURNING recebedor_id"

	err := r.DB.QueryRow(query, recebedor.CpfCnpj, recebedor.Nome, recebedor.TipoChavePix, recebedor.ChavePix, recebedor.Status, recebedor.Email).Scan(&recebedor.Id)
	if err != nil {
		return err
	}
	return nil
}
