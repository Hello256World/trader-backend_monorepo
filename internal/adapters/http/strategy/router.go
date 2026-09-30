package strategy

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine, handler Handlers) {
	if engine == nil || handler == nil {
		return
	}

	v1 := engine.Group("/v1")

	v1.POST("/strategies", handler.Create)
	v1.GET("/strategies/:strategy-id", handler.GetByID)
}
