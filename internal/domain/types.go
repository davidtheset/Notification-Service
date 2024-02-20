package domain

import "errors"

// NotificationRequest define the request to send a notification.
type NotificationRequest struct {
	Recipient        string `json:"recipient"`
	NotificationType string `json:"notification_type"`
}

// NotificationLimiter define the interface to limit the notifications sent.
type NotificationLimiter interface {
	CanSendNotification(recipient string, notificationType string) bool
	RecordNotificationSent(recipient string, notificationType string)
}

// ErrorateLimitExceeded define the error when the rate limit is exceeded.
var ErrRateLimitExceeded = errors.New("rate limit exceeded")
