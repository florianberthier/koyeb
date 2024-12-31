package service

import (
	"fmt"
	"koyeb/models"
	"koyeb/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateService(c *gin.Context) {
	var request models.CreateServiceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.Validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	content, err := utils.Fetch(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch content: %v", err)})
		return
	}

	port, ok := s.Jobs[name]
	if !ok {
		port, err = utils.GetRandomPort(3001, 4000)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting available port: %v", err)})
			return
		}

		s.Jobs[name] = port
	}

	job, err := createNomadJob(name, request.Script, content, port)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating Nomad job: %v", err)})
		return
	}

	_, _, err = s.NomadClient.Jobs().Register(job, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error registering job: %v", err)})
		return
	}

	c.JSON(http.StatusOK, models.CreateServiceResponse{
		URL: fmt.Sprintf("http://127.0.0.1:%d", port),
	})
}

func (s *Service) GetServices(c *gin.Context) {
	response := []models.ServiceResponse{}
	for name, port := range s.Jobs {
		response = append(response, models.ServiceResponse{
			Name: name,
			URL:  fmt.Sprintf("http://127.0.0.1:%d", port),
		})
	}

	c.JSON(http.StatusOK, response)
}
