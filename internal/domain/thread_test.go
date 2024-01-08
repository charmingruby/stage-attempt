package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreadJSONSerialization(t *testing.T) {
	thread := NewThread(
		1,
		1,
		"email dispatch.",
		"sender email is not verified.",
	)

	// Serialization
	data, err := json.Marshal(thread)
	assert.NoError(t, err)

	// Deserialization
	var newThread Thread
	err = json.Unmarshal(data, &newThread)
	assert.NoError(t, err)

	assert.Equal(t, thread.ID, newThread.ID)
	assert.Equal(t, thread.Context, newThread.Context)
	assert.Equal(t, thread.ErrorLog, newThread.ErrorLog)
	assert.Equal(t, thread.SolvedAt, newThread.SolvedAt)
	assert.Equal(t, thread.DeletedAt, newThread.DeletedAt)
}
