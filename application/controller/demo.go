package controller

import (
	"ginweb/code"
	"ginweb/config"
	"ginweb/pkg/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DemoController struct {
	core.Controller
}

func (d DemoController) Demo(ctx *gin.Context) {
	result := core.ResponseData{}
	result.Code = code.Success
	result.Message = code.Message[code.Success]
	config.AppLogger.Sugar().Debugw("info", zap.Any("msg", result))
	d.ResponseOk(ctx, result)
	return
}
