package domain

import "github.com/charmingruby/stage-attempt/helpers"

const (
	notification_pending    = "pending"
	notification_canceled   = "canceled"
	notification_processing = "processing"
	notification_delivering = "delivering"
	notification_delivered  = "delivered"

	payment_pending    = "pending"
	payment_canceled   = "canceled"
	payment_processing = "processing"
	payment_delivering = "delivering"
	payment_delivered  = "delivered"

	threat_pending    = "pending"
	threat_canceled   = "canceled"
	threat_processing = "processing"
	threat_solved     = "solved"

	order_pending          = "pending"
	order_awaiting_payment = "awaiting_payment"
	order_developing       = "developing"
	order_delivered        = "delivered"
	order_maintaince       = "maintaince"
	order_threatened       = "threatened"
)

func PaymentStatus() map[string]string {
	return map[string]string{
		payment_pending:    "pending",
		payment_canceled:   "canceled",
		payment_processing: "processing",
		payment_delivering: "delivering",
		payment_delivered:  "delivered",
	}
}

func NotificationStatus() map[string]string {
	return map[string]string{
		notification_pending:    "pending",
		notification_canceled:   "canceled",
		notification_processing: "processing",
		notification_delivering: "delivering",
		notification_delivered:  "delivered",
	}
}

func ThreatStatus() map[string]string {
	return map[string]string{
		threat_pending:    "pending",
		threat_canceled:   "canceled",
		threat_processing: "processing",
		threat_solved:     "solved",
	}
}

func OrderStatus() map[string]string {
	return map[string]string{
		order_pending:          "pending",
		order_awaiting_payment: "awaiting_payment",
		order_developing:       "developing",
		order_delivered:        "delivered",
		order_maintaince:       "maintaince",
		order_threatened:       "threatened",
	}
}

func IsStatusValid(m map[string]string, v string) bool {
	return helpers.IsItemInStringMap(m, v)
}
