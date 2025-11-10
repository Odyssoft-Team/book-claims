package model

import "github.com/google/uuid"

type ComplaintSequence struct {
	TenantID     uuid.UUID
	Year         int
	CurrentValue int64
}
