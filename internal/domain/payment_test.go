package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentJSONSerialization(t *testing.T) {
	paymentType := PaymentType{
		CreditPayment: &CreditOpt{
			AmountInCents:    100,
			AmountOfPartials: 4,
		},
		CashPayment: nil,
	}

	payment := NewPayment(paymentType, 1, 1)

	// Serialization
	data, err := json.Marshal(payment)
	assert.NoError(t, err)

	// Deserialization
	var newPayment Payment
	err = json.Unmarshal(data, &newPayment)
	assert.NoError(t, err)

	assert.Equal(t, payment.ID, newPayment.ID)
	assert.Equal(t, payment.PayerID, newPayment.PayerID)
	assert.Equal(t, payment.Payment, newPayment.Payment)

}
