package app

import (
	"testing"

	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CriarRecebedor(recebedor *domain.Recebedor) error {
	args := m.Called(recebedor)
	return args.Error(0)
}
func (m *MockRepository) EditarEmailRecebedor(id uint, email string) error {
	args := m.Called(id, email)
	return args.Error(0)
}
func (m *MockRepository) BuscarRecebedorPorId(id uint) (*domain.Recebedor, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Recebedor), args.Error(1)
}

func (m *MockRepository) EditarRecebedor(recebedor *domain.Recebedor) error {
	args := m.Called(recebedor)
	return args.Error(0)
}
func (m *MockRepository) DeletarRecebedor(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockRepository) DeletarRecebedores(ids []uint) error {
	args := m.Called(ids)
	return args.Error(0)
}

func (m *MockRepository) BuscarRecebedoresPorCampo(valor, nomeCampo string, offset int) ([]*domain.Recebedor, error) {
	args := m.Called(valor, nomeCampo, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Recebedor), args.Error(1)
}
func (m *MockRepository) ContarRecebedoresPorCampo(valor, nomeCampo string) (int, error) {
	args := m.Called(valor, nomeCampo)
	if args.Get(0) == nil {
		return 0, args.Error(1)
	}
	return args.Get(0).(int), args.Error(1)
}
func mockLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
func TestDeletarRecebedores_Success(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	ids := []uint{1, 2, 3, 4}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("BuscarRecebedorPorId", uint(2)).Return(recebedor, nil)
	repo.On("BuscarRecebedorPorId", uint(3)).Return(recebedor, nil)
	repo.On("BuscarRecebedorPorId", uint(4)).Return(recebedor, nil)
	repo.On("DeletarRecebedor", uint(1)).Return(nil)
	repo.On("DeletarRecebedor", uint(2)).Return(nil)
	repo.On("DeletarRecebedor", uint(3)).Return(nil)
	repo.On("DeletarRecebedor", uint(4)).Return(nil)
	err := svc.DeletarRecebedores(ids)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestBuscarRecebedorPorNome_Success(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}

	recebedores := []*domain.Recebedor{
		{
			Id:           57,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CPF",
			ChavePix:     "712.468.680-00114",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
		{
			Id:           58,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CNPJ",
			ChavePix:     "71.246.868/0001-14",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
	}
	esperado := &domain.PaginaRecebedores{
		Total:        2,
		PorPagina:    10,
		PaginaAtual:  1,
		TotalPaginas: 1,
		Recebedores:  recebedores,
	}

	nome := "flavio"
	nomeCampo := "nome"
	paginacao := 1
	repo.On("ContarRecebedoresPorCampo", nome, nomeCampo).Return(2, nil)
	repo.On("BuscarRecebedoresPorCampo", nome, nomeCampo, 0).Return(recebedores, nil)

	response, err := svc.BuscarRecebedoresPorNome(nome, paginacao)
	assert.NoError(t, err)
	assert.Equal(t, esperado, response)
	repo.AssertExpectations(t)
}

func TestBuscarRecebedorPorStatus_Success(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}

	recebedores := []*domain.Recebedor{
		{
			Id:           57,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CPF",
			ChavePix:     "712.468.680-00114",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
		{
			Id:           58,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CNPJ",
			ChavePix:     "71.246.868/0001-14",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
	}
	esperado := &domain.PaginaRecebedores{
		Total:        2,
		PorPagina:    10,
		PaginaAtual:  1,
		TotalPaginas: 1,
		Recebedores:  recebedores,
	}

	valorCampo := "Rascunho"
	nomeCampo := "status_recebedor"
	paginacao := 1
	repo.On("ContarRecebedoresPorCampo", valorCampo, nomeCampo).Return(2, nil)
	repo.On("BuscarRecebedoresPorCampo", valorCampo, nomeCampo, 0).Return(recebedores, nil)

	response, err := svc.BuscarRecebedoresPorStatus(valorCampo, paginacao)
	assert.NoError(t, err)
	assert.Equal(t, esperado, response)
	repo.AssertExpectations(t)
}

func TestBuscarRecebedorPorChave_Success(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}

	recebedores := []*domain.Recebedor{
		{
			Id:           57,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CPF",
			ChavePix:     "712.468.680-00114",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
		{
			Id:           58,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CNPJ",
			ChavePix:     "71.246.868/0001-14",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
	}
	esperado := &domain.PaginaRecebedores{
		Total:        2,
		PorPagina:    10,
		PaginaAtual:  1,
		TotalPaginas: 1,
		Recebedores:  recebedores,
	}

	valorCampo := "71.246.868/0001-14"
	nomeCampo := "chave_pix"
	paginacao := 1
	repo.On("ContarRecebedoresPorCampo", valorCampo, nomeCampo).Return(2, nil)
	repo.On("BuscarRecebedoresPorCampo", valorCampo, nomeCampo, 0).Return(recebedores, nil)

	response, err := svc.BuscarRecebedoresPorChave(valorCampo, paginacao)
	assert.NoError(t, err)
	assert.Equal(t, esperado, response)
	repo.AssertExpectations(t)
}

func TestBuscarRecebedorPorChave_ChaveInvalida(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}

	valorCampo := "xxxxxx"
	paginacao := 1
	_, err := svc.BuscarRecebedoresPorChave(valorCampo, paginacao)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrChaveInvalida, err)
}

func TestBuscarRecebedorPorTipoChave_Success(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}

	recebedores := []*domain.Recebedor{
		{
			Id:           57,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CPF",
			ChavePix:     "712.468.680-00114",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
		{
			Id:           58,
			CpfCnpj:      "081.312.395-00",
			Nome:         "flavio",
			TipoChavePix: "CNPJ",
			ChavePix:     "71.246.868/0001-14",
			Status:       "Rascunho",
			Email:        "novo@em.com",
		},
	}
	esperado := &domain.PaginaRecebedores{
		Total:        2,
		PorPagina:    10,
		PaginaAtual:  1,
		TotalPaginas: 1,
		Recebedores:  recebedores,
	}

	valorCampo := "CNPJ"
	nomeCampo := "tipo_chave_pix"
	paginacao := 1
	repo.On("ContarRecebedoresPorCampo", valorCampo, nomeCampo).Return(2, nil)
	repo.On("BuscarRecebedoresPorCampo", valorCampo, nomeCampo, 0).Return(recebedores, nil)

	response, err := svc.BuscarRecebedoresPorTipoChavePix(valorCampo, paginacao)
	assert.NoError(t, err)
	assert.Equal(t, esperado, response)
	repo.AssertExpectations(t)
}
func TestBuscarRecebedorPorTipo_TipoInvalido(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}

	valorCampo := "xxxxxx"
	paginacao := 1
	_, err := svc.BuscarRecebedoresPorTipoChavePix(valorCampo, paginacao)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrTipoChaveInvalida, err)
}

func TestBuscarRecebedor_NaoEncontrado(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	repo.On("BuscarRecebedorPorId", uint(2)).Return(nil, domain.ErrRecebedorNaoEncontrado)
	recebedor, err := svc.BuscarRecebedorById(uint(2))
	assert.Error(t, err)
	assert.Equal(t, domain.ErrRecebedorNaoEncontrado, err)
	assert.Nil(t, recebedor)
	repo.AssertExpectations(t)
}
func TestDeletarRecebedor_Success(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("DeletarRecebedor", uint(1)).Return(nil)
	err := svc.DeletarRecebedor(uint(1))
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
func TestEditarRecebedor_Success(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("EditarRecebedor", recebedor).Return(nil)
	err := svc.EditarRecebedor(recebedor)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestEditarRecebedor_NomeInvalido(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "jj",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("EditarRecebedor", recebedor).Return(nil)
	err := svc.EditarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrNomeInvalido, err)

}
func TestEditarEmailRecebedor_Success(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	email := "flavio@teste.com"
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("EditarEmailRecebedor", uint(1), email).Return(nil)
	err := svc.EditarEmailRecebedor(uint(1), email)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
func TestEditarEmailRecebedor_EmailInvalido(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	email := "flavio@teste"
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("EditarEmailRecebedor", uint(1), email).Return(nil)
	err := svc.EditarEmailRecebedor(uint(1), email)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrEmailInvalido, err)
}
func TestEditarEmailRecebedor_RecebedorNaoEncontrado(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	email := "flavio@teste.com"
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, domain.ErrRecebedorNaoEncontrado)
	repo.On("EditarEmailRecebedor", uint(1), email).Return(nil)
	err := svc.EditarEmailRecebedor(uint(1), email)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrRecebedorNaoEncontrado, err)
}

func TestEditarRecebedor_CpfInvalido(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-62",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("EditarRecebedor", recebedor).Return(nil)
	err := svc.EditarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrCpfInvalido, err)
}
func TestEditarRecebedor_RecebedorNaoEncontrado(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-62",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
	}
	repo.On("BuscarRecebedorPorId", uint(1)).Return(nil, domain.ErrRecebedorNaoEncontrado)
	repo.On("EditarRecebedor", recebedor).Return(domain.ErrRecebedorNaoEncontrado)
	err := svc.EditarRecebedor(recebedor)
	assert.Error(t, domain.ErrRecebedorNaoEncontrado)
	assert.Equal(t, domain.ErrRecebedorNaoEncontrado, err)
}

func TestEditarRecebedor_RecebedorNaoPermiteEdicao(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-62",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
		Status:       "Validado",
	}
	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedor, nil)
	repo.On("EditarRecebedor", recebedor).Return(nil)
	err := svc.EditarRecebedor(recebedor)
	assert.Error(t, domain.ErrRecebedorNaoPermiteEdicao)
	assert.Equal(t, domain.ErrRecebedorNaoPermiteEdicao, err)
}

func TestCreateRecebedor_SuccessoChaveCpf(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "CPF",
		ChavePix:     "515.762.030-69",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateRecebedor_SuccessoChaveAleatoria(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "CHAVE_ALEATORIA",
		ChavePix:     "46892703-d647-4a2c-a6be-a6e0f1488da7",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateRecebedor_SuccessoCnpj(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "41.916.896/0001-30",
		Nome:         "João da Silva",
		TipoChavePix: "TELEFONE",
		ChavePix:     "5579998765676",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
func TestCreateRecebedor_Success(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "CPF",
		ChavePix:     "081.312.395-00",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateRecebedor_SuccessChaveCnpj(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "CNPJ",
		ChavePix:     "41.916.896/0001-30",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateRecebedor_EmailInvalido(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "CPF",
		ChavePix:     "515.762.030-69",
		Email:        "joao@example",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrEmailInvalido, err)

}

func TestCreateRecebedor_ChaveETipoNaoCorresponde(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "CNPJ",
		ChavePix:     "515.762.030-69",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrChaveTipoNaoCorresponde, err)

}

func TestCreateRecebedor_TipoChaveInvalida(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "essa é uma chave inválida",
		ChavePix:     "515.762.030-69",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrTipoChaveInvalida, err)

}
func TestCreateRecebedor_CnpjInvalido(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "41.916.896/0002-30",
		Nome:         "João da Silva",
		TipoChavePix: "TELEFONE",
		ChavePix:     "799965474828",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrCnpjInvalido, err)

}
func TestCreateRecebedor_ChavePixInvalida(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "41.916.896/0001-30",
		Nome:         "João da Silva",
		TipoChavePix: "CHAVE_ALEATORIA",
		ChavePix:     "799965474828",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrChaveInvalida, err)

}

func TestCreateRecebedor_CpfInvalido(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "08122239522",
		Nome:         "João da Silva",
		TipoChavePix: "TELEFONE",
		ChavePix:     "799965474828",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrCpfInvalido, err)

}
func TestCreateRecebedor_TipoChaveTelefoneInvalida(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "41.916.896/0001-30",
		Nome:         "João da Silva",
		TipoChavePix: "TELEFONE",
		ChavePix:     "799965474828",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrChaveInvalida, err)

}

func TestCreateRecebedor_TipoChaveAleatoriaInvalida(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "41.916.896/0001-30",
		Nome:         "João da Silva",
		TipoChavePix: "CHAVE_ALEATORIA",
		ChavePix:     "0f1488da7",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrChaveInvalida, err)

}

func TestCreateRecebedor_TipoChaveCnpjInvalida(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "41.916.896/0001-30",
		Nome:         "João da Silva",
		TipoChavePix: "CPF",
		ChavePix:     "0f1488da7",
		Email:        "joao@example.com",
	}

	repo.On("CriarRecebedor", recebedor).Return(nil)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrChaveInvalida, err)

}
