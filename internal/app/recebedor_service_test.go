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
func mockLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
func TestCreateUser_Success(t *testing.T) {

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

func TestCreateUser_SuccessChaveCnpj(t *testing.T) {

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

func TestCreateUser_EmailInvalido(t *testing.T) {

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
	assert.EqualError(t, err, domain.ErrEmailInvalido.Error())

}

func TestCreateUser_CnpjInvalido(t *testing.T) {

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
	assert.EqualError(t, err, domain.ErrCnpjInvalido.Error())

}
func TestCreateUser_ChavePixInvalida(t *testing.T) {

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
	assert.EqualError(t, err, domain.ErrChaveInvalida.Error())

}

func TestCreateUser_CpfInvalido(t *testing.T) {

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
	assert.EqualError(t, err, domain.ErrCpfInvalido.Error())

}
func TestCreateUser_TipoChaveTelefoneInvalida(t *testing.T) {

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
	assert.EqualError(t, err, domain.ErrChaveInvalida.Error())

}

func TestCreateUser_TipoChaveAleatoriaInvalida(t *testing.T) {

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
	assert.EqualError(t, err, domain.ErrChaveInvalida.Error())

}

func TestCreateUser_TipoChaveCnpjInvalida(t *testing.T) {

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
	assert.EqualError(t, err, domain.ErrChaveInvalida.Error())

}
