package domain

import "time"

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

func (t *Threat) Validate() error {
	// TODO: validation

	return nil
}

func (t *Threat) StatusPending() *Threat {
	status := ThreatStatus()

	t.Status = (status[threat_pending])

	return t
}
