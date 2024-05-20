package http

import (
	"net/http"

	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var status int
			var message string

			switch err {
			case domain.ErrEmailInvalido, domain.ErrCpfInvalido, domain.ErrChaveTipoNaoCorresponde, domain.ErrCnpjInvalido, domain.ErrNomeInvalido, domain.ErrTipoChaveInvalida, domain.ErrChaveInvalida:
				status = http.StatusBadRequest
				message = err.Error()
			case domain.ErrRecebedorNaoEncontrado:
				status = http.StatusNotFound
				message = err.Error()
			case domain.ErrRecebedorNaoPermiteEdicao:
				status = http.StatusConflict
				message = err.Error()
			default:
				status = http.StatusInternalServerError
				message = "erro interno no servidor"
				if _, ok := err.(domain.ErrRecebedoresNaoDeletados); ok {
					status = http.StatusMultiStatus
					message = err.Error()
				}

			}

			c.JSON(status, ErrorResponse{
				Code:    status,
				Message: message,
			})
		}

	}
}
