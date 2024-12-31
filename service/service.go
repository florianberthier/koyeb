package service

import (
	"fmt"
	"io"
	"koyeb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

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

	content, err := fetch(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch content: %v", err)})
		return
	}

	job, err := createNomadJob(name, request.Script, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating Nomad job: %v", err)})
		return
	}

	nomadResp, _, err := s.NomadClient.Jobs().Register(job, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error registering job: %v", err)})
		return
	}

	// wait for the job to be running

	fmt.Printf("Created service: %+v\n", *nomadResp)

	c.JSON(http.StatusOK, models.CreateServiceResponse{
		URL: "http://127.0.0.1:80",
	})
}

func (s *Service) GetAllocations(c *gin.Context) {
	allocs, _, err := s.NomadClient.Allocations().List(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching allocations: %v", err)})
		return
	}
	response := []models.ServiceResponse{}
	for _, alloc := range allocs {
		allocInfo, _, err := s.NomadClient.Allocations().Info(alloc.ID, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error fetching allocation info: %v", err)})
			return
		}

		if len(allocInfo.Resources.Networks) == 0 {
			fmt.Println("No network resources found for allocation, skipping")
			continue
		}

		response = append(response, models.ServiceResponse{
			ID:     alloc.ID,
			URL:    fmt.Sprintf("http://%s:80", allocInfo.Resources.Networks[0].IP),
			Status: allocInfo.ClientStatus,
		})
	}

	c.JSON(http.StatusOK, response)
}
