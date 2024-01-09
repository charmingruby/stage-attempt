package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotificationJSONSerialization(t *testing.T) {
	notification := NewNotification(
		1,
		1,
		1,
		"order status updated.",
	)

	// Serialization
	data, err := json.Marshal(notification)
	assert.NoError(t, err)

	// Deserialization
	var newNotification Notification
	err = json.Unmarshal(data, &newNotification)
	assert.NoError(t, err)

	assert.Equal(t, notification.ID, newNotification.ID)
	assert.Equal(t, notification.Content, newNotification.Content)
	assert.Equal(t, notification.ReceiverID, newNotification.ReceiverID)
	assert.Equal(t, notification.SenderID, newNotification.SenderID)
	assert.Equal(t, notification.Status, newNotification.Status)

}
