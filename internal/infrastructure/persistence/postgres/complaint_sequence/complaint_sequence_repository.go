package complaintsequence

import (
	"claimbook-api/internal/core/port"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type complaintSequencePGRepository struct {
	db *gorm.DB
}

func NewComplaintSequencePGRepository(db *gorm.DB) port.ComplaintSequenceRepository {
	return &complaintSequencePGRepository{db: db}
}

func (r *complaintSequencePGRepository) GenerateCodePublic(ctx context.Context, tenantID uuid.UUID, prefix string) (string, error) {
	year := time.Now().Year()
	var seq int64

	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return "", tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	seqRecord := &ComplaintSequenceModel{
		TenantID:     tenantID,
		Year:         year,
		CurrentValue: 1,
	}

	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "tenant_id"}, {Name: "year"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"current_value": gorm.Expr("complaint_sequence.current_value + 1")}),
	}).Create(seqRecord).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Raw(`SELECT current_value FROM complaint_sequence WHERE tenant_id = ? AND year = ?`, tenantID, year).Scan(&seq).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	code := fmt.Sprintf("%s-%d-%06d", prefix, year, seq)
	return code, nil
}
