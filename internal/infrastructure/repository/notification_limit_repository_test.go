package repository

import (
	"testing"
	"time"
)

// TestCanSendNotification verify that the notification limiter allows or rejects notifications as expected.
func TestCanSendNotification(t *testing.T) {
	// Create a new instance of InMemoryNotificationLimiter
	limiter := NewInMemoryNotificationLimiter()

	// Test cases
	tests := []struct {
		name             string
		prepare          func() // Function to prepare the state for the test
		recipient        string
		notificationType string
		want             bool // Result expected
	}{
		{
			name:             "Allow first status notification",
			prepare:          func() {}, // its not necessary to prepare the state for this test
			recipient:        "user@example.com",
			notificationType: "status",
			want:             true,
		},
		{
			name: "Reject third status notification in the same minute",
			prepare: func() {
				// Simulate two notifications of status sent in the last minute
				limiter.RecordNotificationSent("user@example.com", "status")
				limiter.RecordNotificationSent("user@example.com", "status")
			},
			recipient:        "user@example.com",
			notificationType: "status",
			want:             false,
		},
		{
			name: "Allow second marketing notification at the same time",
			prepare: func() {
				// Simulate one notification of marketing sent in the last hour
				limiter.RecordNotificationSent("user@example.com", "marketing")
			},
			recipient:        "user@example.com",
			notificationType: "marketing",
			want:             true,
		},
		{
			name: "Allow notifications for different recipients",
			prepare: func() {
				limiter.RecordNotificationSent("user1@example.com", "news")
			},
			recipient:        "user2@example.com",
			notificationType: "news",
			want:             true,
		},
		{
			name: "Accept notification just after the time interval passes",
			prepare: func() {
				// Clean the state and set the last notification time to 24 hours ago
				limiter.recipientLastNotification = make(map[string]map[string][]time.Time)

				// Adjust the last notification time to 24 hours ago
				lastNotificationTime := time.Now().Add(-24 * time.Hour)
				limiter.recipientLastNotification["user@example.com"] = map[string][]time.Time{
					"news": {lastNotificationTime},
				}
			},
			recipient:        "user@example.com",
			notificationType: "news",
			want:             true,
		},
		{
			name: "Allow different types of notifications simultaneously",
			prepare: func() {
				limiter.RecordNotificationSent("user@example.com", "news")
			},
			recipient:        "user@example.com",
			notificationType: "status",
			want:             true,
		},
		{
			name: "Reject notification after reaching exact limit",
			prepare: func() {
				for i := 0; i < 2; i++ {
					limiter.RecordNotificationSent("user@example.com", "status")
				}
			},
			recipient:        "user@example.com",
			notificationType: "status",
			want:             false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare the state before each test
			tt.prepare()

			// Call the function CanSendNotification and verify the result
			got := limiter.CanSendNotification(tt.recipient, tt.notificationType)
			if got != tt.want {
				t.Errorf("CanSendNotification() = %v, want %v", got, tt.want)
			}

			// Clean the state after each test
			limiter.recipientLastNotification = make(map[string]map[string][]time.Time)
		})
	}
}
