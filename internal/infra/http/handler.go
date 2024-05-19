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

func (h *RecebedorHandler) BuscarRecebedorPorId(c *gin.Context) {
	idStr := c.Param("id")
	idTmp, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id inválido",
		})
		return
	}
	var id = uint(idTmp)
	recebedor, err := h.service.BuscarRecebedorById(id)
	if err != nil {
		h.logger.Error("consultando recebedor", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recebedor)
}
