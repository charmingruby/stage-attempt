package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionJSONSerialization(t *testing.T) {
	paymentType := PaymentType{
		CreditPayment: &CreditOpt{
			Amount:           10,
			AmountOfPartials: 4,
		},
		CashPayment: nil,
	}

	transaction := NewTransaction(paymentType, 1, 1)

	// Serialization
	data, err := json.Marshal(transaction)
	assert.NoError(t, err)

	// Deserialization
	var newTransaction Transaction
	err = json.Unmarshal(data, &newTransaction)
	assert.NoError(t, err)

	assert.Equal(t, transaction.ID, newTransaction.ID)
	assert.Equal(t, transaction.PayerID, newTransaction.PayerID)
	assert.Equal(t, transaction.Payment, newTransaction.Payment)

}
