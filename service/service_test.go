package service

import (
	"encoding/json"
	"koyeb/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestCreateService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Invalid JSON in request body", func(t *testing.T) {
		s := &Service{Validator: validator.New()}
		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)

		ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("invalid json"))
		ctx.Request.Header.Set("Content-Type", "application/json")

		s.CreateService(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid character")
	})

	t.Run("Validation url error", func(t *testing.T) {
		s := &Service{Validator: validator.New()}
		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)

		requestBody := `{"url":"https://pastebin.com/raw/hEFbnx33","script":false}`
		ctx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		ctx.Request.Header.Set("Content-Type", "application/json")

		s.CreateService(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestGetServices(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := &Service{
		Jobs: map[string]int{
			"service1": 3001,
			"service2": 3002,
		},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	service.GetServices(c)

	assert.Equal(t, http.StatusOK, w.Code)

	response := []models.ServiceResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response, 2)
	assert.Equal(t, true, response[0].Name == "service1" || response[0].Name == "service2")
	assert.Equal(t, true, response[1].Name == "service1" || response[1].Name == "service2")
	assert.NotEqual(t, response[0].Name, response[1].Name)

	assert.Equal(t, true, response[0].URL == "http://127.0.0.1:3001" || response[0].URL == "http://127.0.0.1:3002")
	assert.Equal(t, true, response[1].URL == "http://127.0.0.1:3001" || response[1].URL == "http://127.0.0.1:3002")
	assert.NotEqual(t, response[0].URL, response[1].URL)
}
