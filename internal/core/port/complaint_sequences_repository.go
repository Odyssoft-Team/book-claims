package port

import (
	"context"

	"github.com/google/uuid"
)

type ComplaintSequenceRepository interface {
	GenerateCodePublic(ctx context.Context, tenantID uuid.UUID, prefix string) (string, error)
}
