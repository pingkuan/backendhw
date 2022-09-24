package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pingkuan/backendhw/server/handler"
)

func Route(r *gin.Engine) {
	r.POST("/callback", handler.Callback)
}
