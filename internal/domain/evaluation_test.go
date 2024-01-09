package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluationJSONSerialization(t *testing.T) {
	evaluation := NewEvaluation(
		1,
		1,
		1,
		5,
		"quick delivery, i loved!",
	)

	// Serialization
	data, err := json.Marshal(evaluation)
	assert.NoError(t, err)

	// Deserialization
	var newEvaluation Evaluation
	err = json.Unmarshal(data, &newEvaluation)
	assert.NoError(t, err)

	assert.Equal(t, evaluation.ID, newEvaluation.ID)
	assert.Equal(t, evaluation.Comment, newEvaluation.Comment)
	assert.Equal(t, evaluation.CustomerID, newEvaluation.CustomerID)
	assert.Equal(t, evaluation.OrderID, newEvaluation.OrderID)
	assert.Equal(t, evaluation.Rate, newEvaluation.Rate)

}
