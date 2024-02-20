package application

import (
	"Notification-Service/internal/domain"
	"testing"
)

// mockNotificationLimiter implement a mock of the NotificationLimiter interface.
type mockNotificationLimiter struct {
	canSend bool
}

func (m *mockNotificationLimiter) CanSendNotification(recipient string, notificationType string) bool {
	return m.canSend
}

func (m *mockNotificationLimiter) RecordNotificationSent(recipient string, notificationType string) {
	// Simulate the recording of the notification sent.
}

func TestSendNotification_Allowed(t *testing.T) {
	limiter := &mockNotificationLimiter{canSend: true}
	service := NewNotificationService(limiter)

	request := &domain.NotificationRequest{
		Recipient:        "test@example.com",
		NotificationType: "marketing",
	}

	if err := service.SendNotification(request); err != nil {
		t.Errorf("SendNotification() error = %v, wantErr %v", err, false)
	}
}

func TestSendNotification_NotAllowed(t *testing.T) {
	limiter := &mockNotificationLimiter{canSend: false} // Simulate the rate limit exceeded.
	service := NewNotificationService(limiter)

	request := &domain.NotificationRequest{
		Recipient:        "test@example.com",
		NotificationType: "marketing",
	}

	if err := service.SendNotification(request); err == nil {
		t.Errorf("Expected error when sending notification is not allowed, got nil")
	}
}
