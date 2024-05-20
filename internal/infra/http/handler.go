package http

import (
	"net/http"
	"strconv"

	"github.com/flaviorodolfo/transfeera-challenge/internal/app"
	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type RecebedorHandler struct {
	service *app.RecebedorService
	logger  *zap.Logger
}

type deleteRequest struct {
	Ids []uint `json:"ids"`
}

func formatarErroCampos(validationErrors validator.ValidationErrors) []map[string]string {
	errors := make([]map[string]string, len(validationErrors))
	for i, ve := range validationErrors {
		errors[i] = map[string]string{
			"campo": ve.Field(),
			"error": ve.Tag(),
		}
	}
	return errors
}

func (h *RecebedorHandler) CriarRecebedor(c *gin.Context) {
	var recebedor domain.Recebedor
	if err := c.ShouldBindJSON(&recebedor); err != nil {
		h.logger.Error("Binding json", zap.Error(err))
		c.Error(err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(&recebedor); err != nil {
		h.logger.Error("validação de campos", zap.Error(err))
		campos := formatarErroCampos(err.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "campos obrigatórios",
			"campos":  campos,
		})
		return
	}
	err := h.service.CriarRecebedor(&recebedor)
	if err != nil {
		h.logger.Error("Criando recebedor", zap.Error(err))
		c.Error(err)
		return
	}
	c.Status(http.StatusCreated)
}

func (h *RecebedorHandler) EditarRecebedor(c *gin.Context) {
	var recebedor domain.Recebedor
	if err := c.ShouldBindJSON(&recebedor); err != nil {
		h.logger.Error("Binding json", zap.Error(err))
		c.Error(err)
		return
	}

	err := h.service.EditarRecebedor(&recebedor)
	if err != nil {
		h.logger.Error("editando recebedor", zap.Error(err))
		c.Error(err)
		return
	}
	c.Status(http.StatusCreated)
}

func (h *RecebedorHandler) EditarEmailRecebedor(c *gin.Context) {
	var recebedor domain.Recebedor
	if err := c.ShouldBindJSON(&recebedor); err != nil {
		h.logger.Error("Binding json", zap.Error(err))
		c.Error(err)
		return
	}
	idStr := c.Param("id")
	idTmp, err := strconv.Atoi(idStr)
	if err != nil || idTmp < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id inválido",
		})
		return
	}
	var id = uint(idTmp)
	err = h.service.EditarEmailRecebedor(id, recebedor.Email)
	if err != nil {
		h.logger.Error("editando recebedor", zap.Error(err))
		c.Error(err)
		return
	}
	c.Status(http.StatusCreated)
}

func (h *RecebedorHandler) BuscarRecebedorPorId(c *gin.Context) {
	idStr := c.Param("id")
	idTmp, err := strconv.Atoi(idStr)
	if err != nil || idTmp < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id inválido",
		})
		return
	}
	var id = uint(idTmp)
	recebedor, err := h.service.BuscarRecebedorById(id)
	if err != nil {
		h.logger.Error("consultando recebedor por id", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recebedor)
}

func (h *RecebedorHandler) DeletarRecebedor(c *gin.Context) {
	idStr := c.Param("id")
	idTmp, err := strconv.Atoi(idStr)
	if err != nil || idTmp < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id inválido",
		})
		return
	}
	var id = uint(idTmp)
	err = h.service.DeletarRecebedor(id)
	if err != nil {
		h.logger.Error("deletando recebedor", zap.Error(err))
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *RecebedorHandler) DeletarRecebedores(c *gin.Context) {
	var body deleteRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ids inválidos",
		})
		return
	}
	err := h.service.DeletarRecebedores(body.Ids)
	if err != nil {
		h.logger.Error("deletando recebedores", zap.Error(err))
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *RecebedorHandler) BuscarRecebedorPorNome(c *gin.Context) {
	nome := c.Param("nome")
	pagina, err := strconv.Atoi(c.DefaultQuery("pagina", "1"))
	if err != nil || pagina < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parâmetro de página inválido"})
		return
	}
	recebedor, err := h.service.BuscarRecebedoresPorNome(nome, pagina)
	if err != nil {
		h.logger.Error("consultando recebedor por nome", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recebedor)
}

func (h *RecebedorHandler) BuscarRecebedorPorStatus(c *gin.Context) {
	status := c.Param("status")
	pagina, err := strconv.Atoi(c.DefaultQuery("pagina", "1"))
	if err != nil || pagina < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parâmetro de página inválido"})
		return
	}
	recebedores, err := h.service.BuscarRecebedoresPorStatus(status, pagina)
	if err != nil {
		h.logger.Error("consultando recebedor", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recebedores)
}

func (h *RecebedorHandler) BuscarRecebedorPorChave(c *gin.Context) {
	chave := c.Param("chave")
	pagina, err := strconv.Atoi(c.DefaultQuery("pagina", "1"))
	if err != nil || pagina < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro de página inválido"})
		return
	}
	recebedores, err := h.service.BuscarRecebedoresPorChave(chave, pagina)
	if err != nil {
		h.logger.Error("consultando recebedor por chave", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recebedores)
}

func (h *RecebedorHandler) BuscarRecebedorPorTipoChave(c *gin.Context) {
	tipoChave := c.Param("tipoChave")
	pagina, err := strconv.Atoi(c.DefaultQuery("pagina", "1"))
	if err != nil || pagina < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro de página inválido"})
		return
	}
	recebedores, err := h.service.BuscarRecebedoresPorTipoChavePix(tipoChave, pagina)
	if err != nil {
		h.logger.Error("consultando recebedor por tipo chave", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recebedores)
}
