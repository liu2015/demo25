package jztinfo

import (
	"demo25/model/jztinfo"
	"demo25/model/jztinfo/request"
	"demo25/model/response"
	"demo25/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JztinfoApi struct {
}

var s = service.ServiceApp.Jztinfos

func (a *JztinfoApi) Create(ctx *gin.Context) {
	var info jztinfo.Jztinfo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("失败", ctx)
		return
	}
	ids, _ := uuid.NewUUID()
	info.Info = ids.String()
	if err := s.Create(info); err != nil {
		response.FailWithMessage("失败", ctx)
	} else {
		response.OkWithMessage("成功", ctx)
	}

}

func (a *JztinfoApi) Delete(ctx *gin.Context) {
	id := ctx.Query("ID")
	if err := s.Delete(id); err != nil {
		response.FailWithMessage("失败", ctx)

	} else {
		response.OkWithMessage("成功", ctx)
	}
}

func (a *JztinfoApi) Deletes(ctx *gin.Context) {
	ids := ctx.QueryArray("IDs[]")
	if err := s.Deletes(ids); err != nil {
		response.FailWithMessage("失败", ctx)
	} else {
		response.OkWithMessage("成功", ctx)
	}
}

func (a *JztinfoApi) Update(ctx *gin.Context) {
	var info jztinfo.Jztinfo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("失败", ctx)
		return
	}
	if err := s.Updates(info); err != nil {
		response.FailWithMessage("失败", ctx)
	} else {
		response.OkWithMessage("成功", ctx)
	}
}

func (a *JztinfoApi) Findone(ctx *gin.Context) {
	id := ctx.Query("ID")
	if list, err := s.FindOne(id); err != nil {
		response.FailWithMessage("失败", ctx)
	} else {
		response.OkWithData(gin.H{"jztinfo": list}, ctx)
	}
}

func (a *JztinfoApi) FindList(ctx *gin.Context) {
	var info request.Findjztinfo
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("失败", ctx)
		return
	}
	if list, total, err := s.FindList(info); err != nil {
		response.FailWithMessage("失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "成功", ctx)
	}
}
