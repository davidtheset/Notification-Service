package application

import (
	"Notification-Service/internal/domain"
)

// NotificationService implement a service to send notifications.
type NotificationService struct {
	limiter domain.NotificationLimiter
}
type NotificationSender interface {
	SendNotification(request *domain.NotificationRequest) error
}

// NewNotificationService creates a new instance of NotificationService.
func NewNotificationService(limiter domain.NotificationLimiter) *NotificationService {
	return &NotificationService{
		limiter: limiter,
	}
}

// SendNotification sends a notification to the recipient.
func (s *NotificationService) SendNotification(request *domain.NotificationRequest) error {
	// Verify if the recipient has exceeded the rate limit.
	if !s.limiter.CanSendNotification(request.Recipient, request.NotificationType) {
		return domain.ErrRateLimitExceeded
	}
	// TODO: Implement the logic to send the notification.

	s.limiter.RecordNotificationSent(request.Recipient, request.NotificationType)
	return nil
}
