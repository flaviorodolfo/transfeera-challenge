package app

import (
	"regexp"
	"strings"

	"github.com/flaviorodolfo/transfeera-challenge/internal/app/validator"
	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
	"go.uber.org/zap"
)

const (
	//nome dos campos para auxiliar a busca por coluna da tabela
	campoNome         = "nome"
	campoChavePix     = "chave_pix"
	campoTipoChavePix = "tipo_chave_pix"
	campoStatus       = "status_recebedor"

	porPagina = 10 //valor padrão da paginação

)

type RecebedorService struct {
	repo   domain.RecebedorRepository
	logger *zap.Logger
}

func NewRecebedorService(repo domain.RecebedorRepository, logger *zap.Logger) *RecebedorService {
	return &RecebedorService{repo: repo, logger: logger}
}

// cria um recebedor, retornar erro se algum dos campos é inválido
func (s *RecebedorService) CriarRecebedor(recebedor *domain.Recebedor) error {
	if err := validarUsuario(recebedor); err != nil {
		s.logger.Error("validando recebedor", zap.Error(err))
		return err
	}
	//normalização do nome do usuário e email e cpf/cnpj
	normalizarCampos(recebedor)
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

// cria um recebedor, retornar erro se algum dos campos é inválido ou se
// o recebedor tem status Validado
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
	//normalização do nome do usuário e email e cpf/cnpj
	normalizarCampos(recebedor)
	if err := s.repo.EditarRecebedor(recebedor); err != nil {
		s.logger.Error("editando recebedor", zap.Error(err))
		return err
	}
	s.logger.Info("recebedor editado com sucesso", zap.Uint("recebedor_id", recebedor.Id))
	return nil
}

// retorna uma lista de recebedores de acordo com os parametros informados
// retorna erro em caso de problema na conexão com o repositório
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
	//calculo dos metadados da paginação
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

// retorna uma lista de recebedores com o nome informado e os metadados da paginacao
// ou erro em caso de problema na conexão com o repositório
func (s *RecebedorService) BuscarRecebedoresPorNome(nome string, pagina int) (*domain.PaginaRecebedores, error) {
	recebedores, err := s.buscarRecebedoresPorCampo(strings.ToLower(nome), campoNome, pagina)
	if err != nil {
		s.logger.Error("consulta de recebedores por nome", zap.Error(err))
		return nil, err
	}
	return recebedores, nil
}

// retorna uma lista de recebedores com o status informado e os metadados da paginacao
// ou erro em caso de problema na conexão com o repositório
func (s *RecebedorService) BuscarRecebedoresPorStatus(status string, pagina int) (*domain.PaginaRecebedores, error) {

	recebedores, err := s.buscarRecebedoresPorCampo(status, campoStatus, pagina)
	if err != nil {
		s.logger.Error("consulta de recebedores por status", zap.Error(err))
		return nil, err
	}
	return recebedores, nil
}

// retorna uma lista de recebedores com a chave informada e os metadados da paginacao
// ou erro em caso de problema na conexão com o repositório ou formato de chave inválida
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

// retorna uma lista de recebedores com o tipo de chave informado e os metadados da paginacao
// ou erro em caso de problema na conexão com o repositório ou tipo de chave inválida
func (s *RecebedorService) BuscarRecebedoresPorTipoChavePix(tipoChave string, pagina int) (*domain.PaginaRecebedores, error) {
	tipo := domain.TipoChavePix(tipoChave)
	if !isTipoValido(tipo) {
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

// retorna um recebedor de acordo com o id informado
// ou erro em caso de problema na conexão com o repositório ou recebedor inexistente
func (s *RecebedorService) BuscarRecebedorById(id uint) (*domain.Recebedor, error) {
	recebedor, err := s.repo.BuscarRecebedorPorId(id)
	if err != nil {
		s.logger.Error("consultando recebedor", zap.Error(err))
		return nil, err
	}
	return recebedor, nil

}

// deleta um recebedor de acordo com o id, retorna erro em caso de recebedor nao existente
// ou problema na conexao com o repositorio
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

// deleta um N recebedores de acordo com os ids informados, caso um ou mais ids não
// existam retorna um erro informando quais foram deletados e quais não
// também retorna erro caso ocorra problema na conexao com o repositorio
func (s *RecebedorService) DeletarRecebedores(ids []uint) error {
	var idsSemSucesso []uint
	var idsComSucesso []uint
	// para cada tentativa de delete ocorre o registro do que obteve sucesso e do que não
	for _, id := range ids {
		if err := s.DeletarRecebedor(id); err != nil {
			idsSemSucesso = append(idsSemSucesso, id)
		} else {
			idsComSucesso = append(idsComSucesso, id)
		}
	}
	if idsSemSucesso != nil {
		return domain.ErrRecebedoresNaoDeletados{IdsComSucesso: idsComSucesso, IdsSemSucesso: idsSemSucesso}
	}
	//for
	// err := s.repo.DeletarRecebedores(ids)
	// if err != nil {
	// 	s.logger.Error("deletando recebedores", zap.Error(err))
	// 	return nil
	// }
	return nil

}

// valida os campos de um usuário
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
	if !isTipoValido(recebedor.TipoChavePix) {
		return domain.ErrTipoChaveInvalida
	}
	if !isChavePixValida(recebedor.ChavePix) {
		return domain.ErrChaveInvalida
	}
	if !isTipoChavePixValida(recebedor.ChavePix, recebedor.TipoChavePix) {
		return domain.ErrChaveTipoNaoCorresponde
	}

	return nil
}

// Função para formatar CPF para o padrão XXX.XXX.XXX-XX
func formatarCpf(cpf string) string {
	re := regexp.MustCompile(`(\d{3})(\d{3})(\d{3})(\d{2})`)
	return re.ReplaceAllString(cpf, "$1.$2.$3-$4")
}

// Função para formatar CNPJ para o padrão XX.XXX.XXX/XXXX-XX
func formatarCnpj(cnpj string) string {
	re := regexp.MustCompile(`(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})`)
	return re.ReplaceAllString(cnpj, "$1.$2.$3/$4-$5")
}

// normaliza os campos de nome, email, cpfCnpj ou chave pix caso necessario
func normalizarCampos(recebedor *domain.Recebedor) {

	re := regexp.MustCompile(`[^\d]`)
	//remove todos os nao digitos do cpfCnpj
	recebedor.CpfCnpj = re.ReplaceAllString(recebedor.CpfCnpj, "")
	//se for CPF aplica a máscara CPF
	if len(recebedor.CpfCnpj) > 11 {
		recebedor.CpfCnpj = formatarCnpj(recebedor.CpfCnpj)
	} else {
		//senao a máscara CNPJ
		recebedor.CpfCnpj = formatarCpf(recebedor.CpfCnpj)
	}
	recebedor.Email = strings.ToLower(recebedor.Email)
	recebedor.Nome = strings.ToLower(recebedor.Nome)
	//verifica se a chave é do tipo CPF ou CNPJ e aplica a devida máscara
	if recebedor.TipoChavePix == domain.Cpf {
		recebedor.ChavePix = formatarCpf(re.ReplaceAllString(recebedor.ChavePix, ""))
	} else if recebedor.TipoChavePix == domain.Cnpj {
		recebedor.ChavePix = formatarCnpj(re.ReplaceAllString(recebedor.ChavePix, ""))
	}
}

// remove todos não digitos da string e valida o CPF ou CNPJ
func validarCpfCnpj(cpfCnpj string) error {
	cpfCnpj = regexp.MustCompile(`[^\d]`).ReplaceAllString(cpfCnpj, "")
	if len(cpfCnpj) > 11 {

		if !validator.ValidarCNPJ(cpfCnpj) {
			return domain.ErrCnpjInvalido
		}
	} else {
		if !validator.ValidarCPF(cpfCnpj) {
			return domain.ErrCpfInvalido
		}
	}
	return nil
}

func isNomeValido(nome string) bool {
	return len(nome) > 2
}

// retorna true se é uma chave pix válida
func isChavePixValida(chave string) bool {
	return validator.ValidarCPF(chave) || validator.ValidarCNPJ(chave) || validator.ValidarTelefone(chave) || validator.ValidarEmail(chave) || validator.ValidarChaveAleatoria(chave)
}

// retorna true se a chave é valida pra o tipo de chave pix
func isTipoChavePixValida(chave string, tipoChave domain.TipoChavePix) bool {
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

// retorna true se é um tipo de chave pix válido
func isTipoValido(tipoChave domain.TipoChavePix) bool {
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
