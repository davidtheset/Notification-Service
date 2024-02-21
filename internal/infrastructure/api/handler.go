package api

import (
	"Notification-Service/internal/application"
	"Notification-Service/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes define the routes for the API and the notification sender to use.
func SetupRoutes(r *gin.Engine, notificationSender application.NotificationSender) {
	r.POST("/send-notification", func(c *gin.Context) {
		var request domain.NotificationRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := notificationSender.SendNotification(&request)
		if err != nil {
			if err == domain.ErrRateLimitExceeded {
				c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
	})
}
