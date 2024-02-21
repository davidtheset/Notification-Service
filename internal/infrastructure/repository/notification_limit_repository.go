package repository

import (
	"sync"
	"time"
)

// InMemoryNotificationLimiter implemen the NotificationLimiter interface to limit the notifications sent.
type InMemoryNotificationLimiter struct {
	recipientLastNotification map[string]map[string][]time.Time
	mu                        sync.Mutex
}

// NewInMemoryNotificationLimiter creates a new instance of InMemoryNotificationLimiter.
func NewInMemoryNotificationLimiter() *InMemoryNotificationLimiter {
	return &InMemoryNotificationLimiter{
		recipientLastNotification: make(map[string]map[string][]time.Time),
	}
}

func (limiter *InMemoryNotificationLimiter) CanSendNotification(recipient string, notificationType string) bool {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	notificationTimes, exists := limiter.recipientLastNotification[recipient][notificationType]
	if !exists {
		// In case the recipient has not received this type of notification before, it can be sent.
		return true
	}
	// Get the limit and interval for the notification type.
	limit, interval := getNotificationLimitAndInterval(notificationType)
	count := 0
	for _, timeSent := range notificationTimes {
		if time.Since(timeSent) < interval {
			count++
		}
	}

	// If the count of notifications sent is less than the limit, the notification can be sent.
	return count < limit
}

func getNotificationLimitAndInterval(notificationType string) (int, time.Duration) {
	// Define the limits and intervals for each type of notification.
	var limits = map[string]struct {
		Limit    int
		Interval time.Duration
	}{
		//TODO 1: Add the limits and intervals for other types of notifications.
		"marketing": {Limit: 3, Interval: time.Hour},
		"news":      {Limit: 1, Interval: 24 * time.Hour},
		"status":    {Limit: 2, Interval: time.Minute},
	}

	if limit, ok := limits[notificationType]; ok {
		return limit.Limit, limit.Interval
	}
	// if the notification type is not found, return 0 limit and 0 interval.
	return 0, 0
}

func (limiter *InMemoryNotificationLimiter) RecordNotificationSent(recipient string, notificationType string) {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	// Verify if the recipient has a history of notifications sent.
	if _, exists := limiter.recipientLastNotification[recipient]; !exists {
		limiter.recipientLastNotification[recipient] = make(map[string][]time.Time)
	}

	// 	add the current time to the list of notifications sent.
	limiter.recipientLastNotification[recipient][notificationType] = append(
		limiter.recipientLastNotification[recipient][notificationType],
		time.Now(),
	)
}
