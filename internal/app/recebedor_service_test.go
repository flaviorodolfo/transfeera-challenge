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
func mockLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

func TestBuscarRecebedor_Success(t *testing.T) {
	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}

	recebedorEsperado := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "CPF",
		ChavePix:     "515.762.030-69",
		Email:        "joao@example.com",
	}

	repo.On("BuscarRecebedorPorId", uint(1)).Return(recebedorEsperado, nil)
	recebedor, err := svc.BuscarRecebedorById(uint(1))
	assert.NoError(t, err)
	assert.Equal(t, recebedorEsperado, recebedor)
	repo.AssertExpectations(t)
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
	repo.On("EditarRecebedor", recebedor).Return(domain.ErrCpfInvalido)
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
	repo.On("EditarRecebedor", recebedor).Return(domain.ErrRecebedorNaoPermiteEdicao)
	err := svc.EditarRecebedor(recebedor)
	assert.Error(t, domain.ErrRecebedorNaoPermiteEdicao)
	assert.Equal(t, domain.ErrRecebedorNaoPermiteEdicao, err)
}

func TestCreateRecebedor_Success(t *testing.T) {

	repo := new(MockRepository)
	svc := &RecebedorService{repo: repo, logger: mockLogger()}
	recebedor := &domain.Recebedor{
		Id:           1,
		CpfCnpj:      "515.762.030-69",
		Nome:         "João da Silva",
		TipoChavePix: "EMAIL",
		ChavePix:     "flavio@transfeera.com",
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

	repo.On("CriarRecebedor", recebedor).Return(domain.ErrEmailInvalido)
	err := svc.CriarRecebedor(recebedor)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrEmailInvalido, err)

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

	repo.On("CriarRecebedor", recebedor).Return(domain.ErrCnpjInvalido)
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

	repo.On("CriarRecebedor", recebedor).Return(domain.ErrChaveInvalida)
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

	repo.On("CriarRecebedor", recebedor).Return(domain.ErrCpfInvalido)
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

	repo.On("CriarRecebedor", recebedor).Return(domain.ErrChaveInvalida)
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

	repo.On("CriarRecebedor", recebedor).Return(domain.ErrChaveInvalida)
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
