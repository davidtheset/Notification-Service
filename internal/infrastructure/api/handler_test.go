package api

import (
	"Notification-Service/internal/domain"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockNotificationService implementa la interfaz necesaria para probar el handler sin dependencias externas.
type mockNotificationService struct {
	sendFunc func(request *domain.NotificationRequest) error
}

func (m *mockNotificationService) SendNotification(request *domain.NotificationRequest) error {
	return m.sendFunc(request)
}

func TestSetupRoutes(t *testing.T) {
	// Initialize Gin and configure paths for testing.
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockService := &mockNotificationService{}
	SetupRoutes(router, mockService)
	t.Run("Successful Notification", func(t *testing.T) {
		mockService.sendFunc = func(request *domain.NotificationRequest) error {
			return nil // Simulate a successful notification.
		}

		body := bytes.NewBufferString(`{"recipient":"test@example.com","notificationType":"marketing"}`)
		req := httptest.NewRequest(http.MethodPost, "/send-notification", body)
		req.Header.Add("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		body := bytes.NewBufferString(`this is not json`)
		req := httptest.NewRequest(http.MethodPost, "/send-notification", body)
		req.Header.Add("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("Rate Limit Exceeded", func(t *testing.T) {
		mockService.sendFunc = func(request *domain.NotificationRequest) error {
			return domain.ErrRateLimitExceeded // Simulate exceeding the rate limit.
		}

		body := bytes.NewBufferString(`{"recipient":"test@example.com","notificationType":"marketing"}`)
		req := httptest.NewRequest(http.MethodPost, "/send-notification", body)
		req.Header.Add("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusTooManyRequests, resp.Code)
	})
}
