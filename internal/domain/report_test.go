package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportJSONDeserialization(t *testing.T) {

	report := NewReport(
		1,
		1,
		"n error fixed",
		"error ocurred during database query is now fixed.",
	)

	// Serialization
	data, err := json.Marshal(report)
	assert.NoError(t, err)

	// Deserialization
	var newReport Report
	err = json.Unmarshal(data, &newReport)
	assert.NoError(t, err)

	assert.Equal(t, report.ID, newReport.ID)
	assert.Equal(t, report.Title, newReport.Title)
	assert.Equal(t, report.Description, newReport.Description)
}
