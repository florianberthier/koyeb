package service

import (
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
