package complaint

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type complaintPGRepository struct {
	db *gorm.DB
}

func NewComplaintPGRepository(db *gorm.DB) port.ComplaintRepository {
	return &complaintPGRepository{db: db}
}

func (r *complaintPGRepository) CreateComplaint(ctx context.Context, complaint *model.Complaint) (*model.Complaint, error) {
	dbModel := ComplaintModelFromDomain(complaint)
	if err := r.db.WithContext(ctx).Create(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *complaintPGRepository) GetByPublicCode(ctx context.Context, code string) (*model.Complaint, error) {
	var dbModel ComplaintModel
	if err := r.db.WithContext(ctx).First(&dbModel, "code_public = ?", code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *complaintPGRepository) GetComplaintById(ctx context.Context, id uuid.UUID) (*model.Complaint, error) {
	var dbModel ComplaintModel
	if err := r.db.WithContext(ctx).First(&dbModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *complaintPGRepository) UpdateComplaint(ctx context.Context, cashbox *model.Complaint) (*model.Complaint, error) {
	dbModel := ComplaintModelFromDomain(cashbox)
	if err := r.db.WithContext(ctx).Save(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *complaintPGRepository) GetComplaints(ctx context.Context) ([]*model.Complaint, error) {
	var dbModels []ComplaintModel
	if err := r.db.WithContext(ctx).Find(&dbModels).Error; err != nil {
		return nil, err
	}

	var result []*model.Complaint
	for _, m := range dbModels {
		result = append(result, m.ToDomain())
	}
	return result, nil
}

func (r *complaintPGRepository) GetSummary(ctx context.Context) (model.SummaryReport, error) {
	var summary model.SummaryReport
	var count int64

	if err := r.db.WithContext(ctx).Model(&ComplaintModel{}).Count(&count).Error; err != nil {
		return summary, err
	}
	summary.TotalComplaints = int(count)

	if err := r.db.WithContext(ctx).Model(&ComplaintModel{}).
		Where("is_closed = ?", true).
		Count(&count).Error; err != nil {
		return summary, err
	}
	summary.Resolved = int(count)

	if err := r.db.WithContext(ctx).Model(&ComplaintModel{}).
		Where("is_closed = ?", false).
		Count(&count).Error; err != nil {
		return summary, err
	}
	summary.Pending = int(count)

	summary.SlaCompliance = "N/A" // placeholder

	return summary, nil
}

func (r *complaintPGRepository) GenerateCodePublic(ctx context.Context, tenantID uuid.UUID, prefix string) (string, error) {
	var seq int64
	err := r.db.WithContext(ctx).Raw("SELECT next_complaint_code(?)", tenantID).Scan(&seq).Error
	if err != nil {
		return "", err
	}
	year := time.Now().Year()
	code := fmt.Sprintf("%s-%d-%06d", prefix, year, seq)
	return code, nil
}
