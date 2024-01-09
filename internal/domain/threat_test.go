package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreatJSONSerialization(t *testing.T) {
	threat := NewThreat(
		1,
		1,
		"email dispatch.",
		"sender email is not verified.",
	)

	// Serialization
	data, err := json.Marshal(threat)
	assert.NoError(t, err)

	// Deserialization
	var newThreat Threat
	err = json.Unmarshal(data, &newThreat)
	assert.NoError(t, err)

	assert.Equal(t, threat.ID, newThreat.ID)
	assert.Equal(t, threat.Context, newThreat.Context)
	assert.Equal(t, threat.ErrorLog, newThreat.ErrorLog)
	assert.Equal(t, threat.Status, newThreat.Status)
	assert.Equal(t, threat.SolvedAt, newThreat.SolvedAt)
	assert.Equal(t, threat.DeletedAt, newThreat.DeletedAt)

}
