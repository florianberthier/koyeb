package server

import (
	"koyeb/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	s := service.Setup()

	r.PUT("/services/:name", s.CreateService)
	r.POST("/services/:name", s.CreateService)

	r.GET("/allocations", s.GetAllocations)
}
