package app

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	server := gin.Default()

	AddRoutes(server)

	server.Run(":8081")

}
