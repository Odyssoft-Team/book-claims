package complaintsequence

import (
	"claimbook-api/internal/core/domain/model"

	"github.com/google/uuid"
)

type ComplaintSequenceModel struct {
	TenantID     uuid.UUID `gorm:"column:tenant_id;primaryKey"`
	Year         int       `gorm:"column:year;primaryKey"`
	CurrentValue int64     `gorm:"column:current_value"`
}

func (u *ComplaintSequenceModel) ToDomain() *model.ComplaintSequence {
	return &model.ComplaintSequence{
		TenantID:     u.TenantID,
		Year:         u.Year,
		CurrentValue: u.CurrentValue,
	}
}
func ComplaintSequenceModelFromDomain(u *model.ComplaintSequence) *ComplaintSequenceModel {
	return &ComplaintSequenceModel{
		TenantID:     u.TenantID,
		Year:         u.Year,
		CurrentValue: u.CurrentValue,
	}
}
func (ComplaintSequenceModel) TableName() string {
	return "complaint_sequence"
}
