package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/pingkuan/backendhw/server/db"
)

func Server(port int) {
	r := gin.Default()

	db.ConnectDB()
	r.Run(fmt.Sprintf(":%d", port))
}
