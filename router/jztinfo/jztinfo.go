package jztinfo

import (
	"demo25/api"

	"github.com/gin-gonic/gin"
)

type JztinfoRouter struct {
}

func (r *JztinfoRouter) Routerinit(Router *gin.RouterGroup) {
	info := Router.Group("v2")
	var a = api.ApiApp.Jztinfos.JztinfoApi
	{
		info.POST("cretae", a.Create)
		info.DELETE("delete", a.Delete)
		info.DELETE("deletes", a.Deletes)
		info.PUT("updates", a.Update)
		info.GET("getone", a.Findone)
		info.GET("getlist", a.FindList)
	}
}
