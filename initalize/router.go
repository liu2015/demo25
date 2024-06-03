package initalize

import (
	"demo25/router"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	def := gin.Default()
	r := def.Group("api")
	routerd := router.RouterApp.Jztinfo.JztinfoRouter
	routerd.Routerinit(r)
	return def
}
