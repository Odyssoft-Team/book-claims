package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
	"time"
)

func CreateComplaintDTOToDomain(c dto.CreateComplaintDTO) *model.Complaint {
	return &model.Complaint{
		TenantID:        c.TenantID,
		LocationID:      c.LocationID,
		Type:            c.Type,
		Status:          c.Status,
		CategoryID:      c.CategoryID,
		SourceID:        c.SourceID,
		ApiKeyID:        c.ApiKeyID,
		CodePublic:      c.CodePublic,
		Description:     c.Description,
		RequestedAction: c.RequestedAction,
		IsClosed:        c.IsClosed,
	}
}

func ComplaintToResponseDTO(complaint *model.Complaint) dto.ComplaintResponse {
	var resolvedAt *time.Time
	if complaint.ResolvedAt != nil {
		resolvedAt = complaint.ResolvedAt
	}
	return dto.ComplaintResponse{
		ID:              complaint.ID,
		TenantID:        complaint.TenantID,
		LocationID:      complaint.LocationID,
		Type:            complaint.Type,
		Status:          complaint.Status,
		CategoryID:      complaint.CategoryID,
		SourceID:        complaint.SourceID,
		ApiKeyID:        complaint.ApiKeyID,
		CodePublic:      complaint.CodePublic,
		Description:     complaint.Description,
		RequestedAction: complaint.RequestedAction,
		CreatedAt:       complaint.CreatedAt,
		UpdatedAt:       *complaint.UpdatedAt,
		ResolvedAt:      resolvedAt,
		IsClosed:        complaint.IsClosed,
	}
}

func UpdateComplaintFromDTO(existing *model.Complaint, dto dto.UpdateComplaintDTO) {
	existing.Status = dto.Status
}
