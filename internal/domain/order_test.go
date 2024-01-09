package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderJSONSerialization(t *testing.T) {
	now := time.Now()

	order := NewOrder(
		1,
		1,
		1,
		"Stage Attempt",
		"Simple API to implement in a Software House system.",
		"API",
		now,
	)

	// Serialization
	data, err := json.Marshal(order)
	assert.NoError(t, err)

	// Deserialization
	var newOrder Order
	err = json.Unmarshal(data, &newOrder)
	assert.NoError(t, err)

	isTheSameDate := order.Deadline.Equal(*newOrder.Deadline)

	assert.Equal(t, order.ID, newOrder.ID)
	assert.Equal(t, order.Name, newOrder.Name)
	assert.Equal(t, order.Description, newOrder.Description)
	assert.True(t, isTheSameDate)
	assert.Equal(t, order.ProductType, newOrder.ProductType)
	assert.Equal(t, order.PaymentID, newOrder.PaymentID)
	assert.Equal(t, order.CustomerID, newOrder.CustomerID)
}
