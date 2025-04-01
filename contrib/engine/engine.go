package engine

import "github.com/gin-gonic/gin"

func NewEngine(mode string, isDefault bool) *gin.Engine {
	gin.SetMode(mode)
	var engine *gin.Engine
	if isDefault {
		engine = gin.Default()
	} else {
		engine = gin.New()
	}
	return engine
}
