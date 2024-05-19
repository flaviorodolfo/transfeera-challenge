package app

import (
	"strings"

	"github.com/flaviorodolfo/transfeera-challenge/internal/app/validator"
	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
	"go.uber.org/zap"
)

const (
	campoNome         = "nome"
	campoChavePix     = "chave_pix"
	campoTipoChavePix = "tipo_chave_pix"
	campoStatus       = "status_recebedor"
	porPagina         = 10
	statusValidando   = "Validando"
	statusRascunho    = "Rascunho"
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

func (s *RecebedorService) buscarRecebedoresPorCampo(nome, nomeDoCampo string, pagina int) (*domain.PaginaRecebedores, error) {
	totalRegistros, err := s.repo.ContarRecebedoresPorCampo(nome, nomeDoCampo)
	if err != nil {
		s.logger.Error("consulta quantidade de registro de recebedores", zap.Error(err))
		return nil, err
	}
	recebedores, err := s.repo.BuscarRecebedoresPorCampo(nome, nomeDoCampo, (pagina-1)*10)
	if err != nil {
		s.logger.Error("consulta de recebedores", zap.Error(err))
		return nil, err
	}
	totalPaginas := totalRegistros / porPagina
	if resto := totalRegistros % porPagina; resto != 0 {
		totalPaginas++
	}
	return &domain.PaginaRecebedores{
		Total:        totalRegistros,
		PorPagina:    porPagina,
		PaginaAtual:  pagina,
		TotalPaginas: totalPaginas,
		Recebedores:  recebedores,
	}, nil
}
func (s *RecebedorService) BuscarRecebedoresPorNome(nome string, pagina int) (*domain.PaginaRecebedores, error) {
	recebedores, err := s.buscarRecebedoresPorCampo(strings.ToLower(nome), campoNome, pagina)
	if err != nil {
		s.logger.Error("consulta de recebedores por nome", zap.Error(err))
		return nil, err
	}
	return recebedores, nil
}

func (s *RecebedorService) BuscarRecebedoresPorStatus(status string, pagina int) (*domain.PaginaRecebedores, error) {

	recebedores, err := s.buscarRecebedoresPorCampo(status, campoStatus, pagina)
	if err != nil {
		s.logger.Error("consulta de recebedores por status", zap.Error(err))
		return nil, err
	}
	return recebedores, nil
}
func (s *RecebedorService) BuscarRecebedoresPorChave(chave string, pagina int) (*domain.PaginaRecebedores, error) {
	if !isChavePixValida(chave) {
		return nil, domain.ErrChaveInvalida
	}
	recebedores, err := s.buscarRecebedoresPorCampo(chave, campoChavePix, pagina)
	if err != nil {
		s.logger.Error("consulta de recebedores por chave", zap.Error(err))
		return nil, err
	}
	return recebedores, nil
}
func (s *RecebedorService) BuscarRecebedoresPorTipoChavePix(tipoChave string, pagina int) (*domain.PaginaRecebedores, error) {
	tp := domain.TipoChavePix(tipoChave)
	if !isTipoChavePixValida(tp) {
		return nil, domain.ErrTipoChaveInvalida
	}
	recebedores, err := s.buscarRecebedoresPorCampo(tipoChave, campoTipoChavePix, pagina)
	if err != nil {
		s.logger.Error("consulta de recebedores por tipo chave pix", zap.Error(err))
		return nil, err
	}
	return recebedores, nil
}
func (s *RecebedorService) EditarEmailRecebedor(id uint, email string) error {

	if !validator.ValidarEmail(email) {
		s.logger.Info("email inválido", zap.String("email", email))
		return domain.ErrEmailInvalido
	}
	_, err := s.repo.BuscarRecebedorPorId(id)
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

func (s *RecebedorService) DeletarRecebedor(id uint) error {
	if _, err := s.BuscarRecebedorById(id); err != nil {
		return err
	}
	err := s.repo.DeletarRecebedor(id)
	if err != nil {
		s.logger.Error("deletando recebedor", zap.Error(err))
		return nil
	}
	return nil

}
func (s *RecebedorService) DeletarRecebedores(ids []uint) error {
	err := s.repo.DeletarRecebedores(ids)
	if err != nil {
		s.logger.Error("deletando recebedores", zap.Error(err))
		return nil
	}
	return nil

}
func validarUsuario(recebedor *domain.Recebedor) error {
	if !isNomeValido(recebedor.Nome) {
		return domain.ErrNomeInvalido
	}
	if recebedor.Email != "" {
		if !validator.ValidarEmail(recebedor.Email) {
			return domain.ErrEmailInvalido
		}
	}
	if err := validarCpfCnpj(recebedor.CpfCnpj); err != nil {
		return err
	}
	if !isTipoChavePixValida(recebedor.TipoChavePix) {
		return domain.ErrTipoChaveInvalida
	}
	if !isChavePixValida(recebedor.ChavePix) {
		return domain.ErrChaveInvalida
	}
	//normalização do nome do usuário e email
	recebedor.Email = strings.ToLower(recebedor.Email)
	recebedor.Nome = strings.ToLower(recebedor.Nome)
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
func isNomeValido(nome string) bool {
	return len(nome) > 2
}
func isChavePixValida(chave string) bool {
	return validator.ValidarCPF(chave) || validator.ValidarCNPJ(chave) || validator.ValidarTelefone(chave) || validator.ValidarEmail(chave) || validator.ValidarChaveAleatoria(chave)
}
func isTipoChavePixValida(tipoChave domain.TipoChavePix) bool {
	switch tipoChave {
	case domain.Cpf:
		return true
	case domain.Cnpj:
		return true
	case domain.Telefone:
		return true
	case domain.Email:
		return true
	case domain.ChaveAleatoria:
		return true
	default:
		return false
	}
}
