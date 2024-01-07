package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserJSONSerialization(t *testing.T) {
	user := NewUser(
		1,
		time.Date(2003, time.May, 31, 0, 0, 0, 0, time.UTC),
		"John",
		"Doe",
		"john@example.com",
		"",
		"image.png",
		"12345678",
	)

	// Serialization
	data, err := json.Marshal(user)
	assert.NoError(t, err)

	// Deserialization
	var newUser User
	err = json.Unmarshal(data, &newUser)
	assert.NoError(t, err)

	assert.Equal(t, user.ID, newUser.ID)
	assert.Equal(t, user.Name, newUser.Name)
	assert.Equal(t, user.LastName, newUser.LastName)
	assert.Equal(t, user.Email, newUser.Email)
	assert.Equal(t, user.Birthdate, newUser.Birthdate)
	assert.Equal(t, user.AvatarUrl, newUser.AvatarUrl)
	assert.Equal(t, user.Role, newUser.Role)
	assert.Equal(t, user.Password, newUser.Password)
}
