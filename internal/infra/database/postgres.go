package database

import (
	"database/sql"
	"fmt"

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

func (r *postgresRecebedorRepository) BuscarRecebedorPorId(id uint) (*domain.Recebedor, error) {
	query := "SELECT recebedor_id,cpf_cnpj, nome, tipo_chave_pix, chave_pix, status_recebedor, email FROM pagamento.recebedores WHERE recebedor_id = $1"

	var recebedor domain.Recebedor
	err := r.DB.QueryRow(query, id).Scan(&recebedor.Id, &recebedor.CpfCnpj, &recebedor.Nome, &recebedor.TipoChavePix, &recebedor.ChavePix, &recebedor.Status, &recebedor.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrRecebedorNaoEncontrado
		}
		return nil, err
	}
	return &recebedor, nil
}

func (r *postgresRecebedorRepository) ContarRecebedoresPorCampo(valor, nomeCampo string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(recebedor_id) FROM pagamento.recebedores WHERE %s = $1", nomeCampo)
	var totalRegistros int
	err := r.DB.QueryRow(query, valor).Scan(&totalRegistros)
	if err != nil {
		return 0, err
	}
	return totalRegistros, nil
}

func (r *postgresRecebedorRepository) BuscarRecebedoresPorCampo(valor, nomeCampo string, offset int) ([]*domain.Recebedor, error) {
	query := fmt.Sprintf("SELECT recebedor_id,cpf_cnpj, nome, tipo_chave_pix, chave_pix, status_recebedor, email FROM pagamento.recebedores WHERE %s = $1 LIMIT 10 OFFSET $2", nomeCampo)
	rows, err := r.DB.Query(query, valor, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrRecebedorNaoEncontrado
		}
		return nil, err
	}
	recebedores := []*domain.Recebedor{}
	for rows.Next() {
		var recebedor domain.Recebedor
		if err := rows.Scan(&recebedor.Id, &recebedor.CpfCnpj, &recebedor.Nome, &recebedor.TipoChavePix, &recebedor.ChavePix, &recebedor.Status, &recebedor.Email); err != nil {
			return nil, err
		}
		recebedores = append(recebedores, &recebedor)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	return recebedores, nil
}

func (r *postgresRecebedorRepository) EditarRecebedor(recebedor *domain.Recebedor) error {
	query := "UPDATE pagamento.recebedores SET "
	values := []interface{}{}
	index := 1
	if recebedor.CpfCnpj != "" {
		query += fmt.Sprintf("cpf_cnpj = $%d, ", index)
		values = append(values, recebedor.CpfCnpj)
		index++
	}
	if recebedor.Nome != "" {
		query += fmt.Sprintf("nome = $%d, ", index)
		values = append(values, recebedor.Nome)
		index++
	}
	if recebedor.TipoChavePix != "" {
		query += fmt.Sprintf("tipo_chave_pix = $%d, ", index)
		values = append(values, recebedor.TipoChavePix)
		index++
	}
	if recebedor.ChavePix != "" {
		query += fmt.Sprintf("chave_pix = $%d, ", index)
		values = append(values, recebedor.ChavePix)
		index++
	}
	if recebedor.Email != "" {
		query += fmt.Sprintf("email = $%d, ", index)
		values = append(values, recebedor.Email)
		index++
	}
	// Remove a vírgula extra
	query = query[:len(query)-2]
	query += " WHERE recebedor_id = $"
	query += fmt.Sprintf("%d", index)
	values = append(values, recebedor.Id)
	fmt.Println(query)
	fmt.Println(values...)
	_, err := r.DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresRecebedorRepository) EditarEmailRecebedor(id uint, email string) error {
	query := "UPDATE pagamento.recebedores SET email = $1 WHERE recebedor_id = $2"
	_, err := r.DB.Exec(query, email, id)
	if err != nil {
		return err
	}
	return nil
}
