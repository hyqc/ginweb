package controller

import (
	"ginweb/code"
	"ginweb/pkg/base"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
	base.Controller
}

func (d DemoController) Demo(ctx *gin.Context) {
	result := base.ResponseData{}
	result.Code = code.Success
	result.Message = code.Message[code.Success]
	d.ResponseOk(ctx, result)
	return
}
