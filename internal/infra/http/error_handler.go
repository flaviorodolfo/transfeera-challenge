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
			case domain.ErrEmailInvalido, domain.ErrCpfInvalido, domain.ErrCnpjInvalido, domain.ErrChaveInvalida:
				status = http.StatusBadRequest
				message = err.Error()
			default:
				status = http.StatusInternalServerError
				message = "erro interno no servidor"
			}

			c.JSON(status, ErrorResponse{
				Code:    status,
				Message: message,
			})
		}

	}
}
