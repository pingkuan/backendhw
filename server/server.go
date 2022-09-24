package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pingkuan/backendhw/server/db"
	"github.com/pingkuan/backendhw/server/routes"
)

func Server(port int) {
	r := gin.Default()

	db.ConnectDB()

	routes.Route(r)

	r.Run(fmt.Sprintf(":%d", port))
}
