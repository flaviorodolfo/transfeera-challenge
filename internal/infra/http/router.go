package http

import (
	"github.com/flaviorodolfo/transfeera-challenge/internal/app"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(service *app.RecebedorService, logger *zap.Logger) *gin.Engine {
	router := gin.Default()
	handler := &RecebedorHandler{service: service, logger: logger}
	router.Use(ErrorHandler())
	v1 := router.Group("/api/v1")
	{
		v1.GET("/recebedores/id/:id", handler.BuscarRecebedorPorId)
		v1.GET("/recebedores/nome/:nome", handler.BuscarRecebedorPorNome)
		v1.GET("/recebedores/status/:status", handler.BuscarRecebedorPorStatus)
		v1.GET("/recebedores/chave/:chave", handler.BuscarRecebedorPorChave)
		v1.GET("/recebedores/tipoChave/:tipoChave", handler.BuscarRecebedorPorTipoChave)
		v1.POST("/recebedores", handler.CriarRecebedor)
		v1.PATCH("/recebedores", handler.EditarRecebedor)
		v1.PATCH("/recebedores/:id", handler.EditarEmailRecebedor)

	}

	return router

}
