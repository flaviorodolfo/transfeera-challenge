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

func (s *RecebedorService) EditarRecebedor(recebedor *domain.Recebedor) error {
	oldRecebedor, err := s.BuscarRecebedorById(recebedor.Id)
	if err != nil {
		s.logger.Error("consultando recebedor", zap.Error(err))
		return err
	}
	if oldRecebedor.Status == "Validado" {
		return domain.ErrRecebedorNaoPermiteEdicao

	}
	if err := validarUsuario(recebedor); err != nil {
		s.logger.Error("validando recebedor", zap.Error(err))
		return err
	}
	if err := s.repo.EditarRecebedor(recebedor); err != nil {
		s.logger.Error("editando recebedor", zap.Error(err))
		return err
	}
	s.logger.Info("recebedor editado com sucesso", zap.Uint("recebedor_id", recebedor.Id))
	return nil
}

func (s *RecebedorService) EditarEmailRecebedor(id uint, email string) error {
	_, err := s.repo.BuscarRecebedorPorId(id)
	if !validator.ValidarEmail(email) {
		s.logger.Info("email inválido", zap.String("email", email))
		return domain.ErrEmailInvalido
	}
	if err != nil {
		s.logger.Error("consultando recebedor", zap.Error(err))
		return err
	}
	err = s.repo.EditarEmailRecebedor(id, email)
	if err != nil {
		s.logger.Error("atualizando email recebedor", zap.Error(err))
		return err

	}
	return nil
}

func (s *RecebedorService) BuscarRecebedorById(id uint) (*domain.Recebedor, error) {
	recebedor, err := s.repo.BuscarRecebedorPorId(id)
	if err != nil {
		s.logger.Error("consultando recebedor", zap.Error(err))
		return nil, err
	}
	return recebedor, nil

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
