package port

import (
	"claimbook-api/internal/core/domain/model"
	"context"

	"github.com/google/uuid"
)

type ComplaintRepository interface {
	CreateComplaint(ctx context.Context, complaint *model.Complaint) (*model.Complaint, error)
	GetByPublicCode(ctx context.Context, code string) (*model.Complaint, error)
	GetComplaintById(ctx context.Context, id uuid.UUID) (*model.Complaint, error)
	UpdateComplaint(ctx context.Context, cashbox *model.Complaint) (*model.Complaint, error)
	GetComplaints(ctx context.Context) ([]*model.Complaint, error)
	GetSummary(ctx context.Context) (model.SummaryReport, error)
}
