package main

import (
	"koyeb/server"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(gin.Recovery())
	server.SetupRouter(r)

	r.Run(":3000")
}
