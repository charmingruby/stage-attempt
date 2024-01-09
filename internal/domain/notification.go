package domain

import "time"

func NewNotification(id, receiverId, senderId int, content string) *Notification {
	now := time.Now()

	n := &Notification{
		ID:         id,
		ReceiverID: receiverId,
		SenderID:   senderId,
		Content:    content,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  nil,
	}

	return n
}

type Notification struct {
	ID         int        `json:"id" db:"id"`
	ReceiverID int        `json:"receiver_id" db:"receiver_id"`
	SenderID   int        `json:"sender_id" db:"sender_id"`
	Content    string     `json:"content" db:"content"`
	Status     string     `json:"status" db:"status"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (n *Notification) StatusPending() *Notification {
	status := NotificationStatus()

	n.Status = status[notification_pending]

	return n
}

func (n *Notification) Validate() error {
	// TODO: validation

	return nil
}
