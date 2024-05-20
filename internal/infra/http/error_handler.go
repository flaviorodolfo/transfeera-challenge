package http

import (
	"encoding/json"
	"net/http"

	"github.com/flaviorodolfo/transfeera-challenge/internal/domain"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func errToJson(erro domain.ErrRecebedoresNaoDeletados) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(erro)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// trata os erros retornados pela camada de servico da aplicação
// insere a mensagem de erro e o status http referente no caso de erros conhecidos
// se é um error inesperado retorna http 500 status code
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var status int
			var message interface{}

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
					if jsonMsg, err := errToJson(err.(domain.ErrRecebedoresNaoDeletados)); err != nil {
						message = err.Error()
					} else {
						message = jsonMsg
					}
				}

			}

			c.JSON(status, ErrorResponse{
				Code:    status,
				Message: message,
			})
		}

	}
}
