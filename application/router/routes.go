package router

import (
	"ginweb/application/controller"
	"github.com/gin-gonic/gin"
)

var (
	demoController = controller.DemoController{}
)

func Routes(e *gin.Engine) {
	r := e.Group("/api")
	{
		r.GET("/demo", demoController.Demo)
	}
}
