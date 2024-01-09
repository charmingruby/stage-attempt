package domain

import "time"

const (
	threat_pending    = "pending"
	threat_canceled   = "canceled"
	threat_processing = "processing"
	threat_solved     = "solved"
)

type Threat struct {
	ID        int        `json:"id" db:"id"`
	TargetID  int        `json:"target_id" db:"target_id"`
	Status    string     `json:"status" db:"status"`
	Context   string     `json:"context" db:"context"`
	ErrorLog  string     `json:"error_log" db:"error_log"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	SolvedAt  *time.Time `json:"solved_at" db:"solved_at"`
}

func NewThreat(id, targetId int, context, errorLog string) *Threat {
	now := time.Now()

	t := &Threat{
		ID:        id,
		TargetID:  targetId,
		Context:   context,
		ErrorLog:  errorLog,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
		SolvedAt:  nil,
	}

	return t
}

func (t *Threat) Validate() error {
	// TODO: validation

	return nil
}

func (t *Threat) StatusPending() *Threat {
	status := threatStatus()

	t.Status = (status[threat_pending])

	return t
}

func threatStatus() map[string]string {
	return map[string]string{
		threat_pending:    "pending",
		threat_canceled:   "canceled",
		threat_processing: "processing",
		threat_solved:     "solved",
	}
}
