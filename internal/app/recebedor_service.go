package app

import (
	"github.com/flaviorodolfo/transfeera-challenge/internal/app/validator"
	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
	"go.uber.org/zap"
)

type RecebedorService struct {
	repo   domain.RecebedorRepository
	logger *zap.Logger
}

func NewRecebedorService(repo domain.RecebedorRepository, logger *zap.Logger) *RecebedorService {
	return &RecebedorService{repo: repo, logger: logger}
}

func (s *RecebedorService) CriarRecebedor(recebedor *domain.Recebedor) error {
	if err := validarUsuario(recebedor); err != nil {
		s.logger.Error("validando recebedor", zap.Error(err))
		return err
	}
	//por definição o status do recebedor no cadastro é Rascunho.
	recebedor.Status = "Rascunho"
	err := s.repo.CriarRecebedor(recebedor)
	if err != nil {
		s.logger.Error("salvando recebedor", zap.Error(err))
		return err
	}
	s.logger.Info("Recebedor criado com sucesso", zap.Uint("ID", recebedor.Id))
	return nil
}

func validarUsuario(recebedor *domain.Recebedor) error {
	if recebedor.Email != "" {
		if !validator.ValidarEmail(recebedor.Email) {
			return domain.ErrEmailInvalido
		}
	}
	if err := validarCpfCnpj(recebedor.CpfCnpj); err != nil {
		return err
	}
	if !isChavePixValida(recebedor.ChavePix, recebedor.TipoChavePix) {
		return domain.ErrChaveInvalida
	}
	return nil
}

func validarCpfCnpj(cpfCnpj string) error {
	if len(cpfCnpj) < 15 {

		if !validator.ValidarCPF(cpfCnpj) {
			return domain.ErrCpfInvalido
		}
	} else {
		if !validator.ValidarCNPJ(cpfCnpj) {
			return domain.ErrCnpjInvalido
		}
	}
	return nil
}

func isChavePixValida(chave string, tipoChave domain.TipoChavePix) bool {
	switch tipoChave {
	case domain.Cpf:
		return validator.ValidarCPF(chave)
	case domain.Cnpj:
		return validator.ValidarCNPJ(chave)
	case domain.Telefone:
		return validator.ValidarTelefone(chave)
	case domain.Email:
		return validator.ValidarEmail(chave)
	case domain.ChaveAleatoria:
		return validator.ValidarChaveAleatoria(chave)
	default:
		return false
	}
}
