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
		v1.GET("/recebedores/:id", handler.BuscarRecebedorPorId)
		v1.POST("/recebedores", handler.CriarRecebedor)
		v1.PATCH("/recebedores", handler.EditarRecebedor)

	}

	return router

}
