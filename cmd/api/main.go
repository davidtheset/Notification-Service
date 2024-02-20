package main

import (
	"Notification-Service/internal/application"
	"Notification-Service/internal/infrastructure/api"
	"Notification-Service/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initilization of the notification service
	notificationLimiter := repository.NewInMemoryNotificationLimiter()
	notificationService := application.NewNotificationService(notificationLimiter)

	// Configuring the routes
	api.SetupRoutes(r, notificationService)

	r.Run(":8080")
}
