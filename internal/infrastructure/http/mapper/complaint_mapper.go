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
		Source:          c.Source,
		ApiKeyID:        c.ApiKeyID,
		CodePublic:      c.CodePublic,
		Description:     c.Description,
		RequestedAction: c.RequestedAction,
		IsClosed:        c.IsClosed,
	}
}

func ComplaintToResponseDTO(complaint *model.Complaint) dto.ComplaintResponse {
	var resolvedAtStr string
	if complaint.ResolvedAt != nil && !complaint.ResolvedAt.IsZero() {
		resolvedAtStr = complaint.ResolvedAt.UTC().Format(time.RFC3339)
	} else {
		resolvedAtStr = ""
	}
	return dto.ComplaintResponse{
		ID:              complaint.ID,
		TenantID:        complaint.TenantID,
		LocationID:      complaint.LocationID,
		Type:            complaint.Type,
		Status:          complaint.Status,
		CategoryID:      complaint.CategoryID,
		Source:          complaint.Source,
		ApiKeyID:        complaint.ApiKeyID,
		CodePublic:      complaint.CodePublic,
		Description:     complaint.Description,
		RequestedAction: complaint.RequestedAction,
		CreatedAt:       complaint.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:       complaint.UpdatedAt.UTC().Format(time.RFC3339),
		ResolvedAt:      resolvedAtStr,
		IsClosed:        complaint.IsClosed,
	}
}

func UpdateComplaintFromDTO(existing *model.Complaint, dto dto.UpdateComplaintDTO) {
	existing.Status = dto.Status
}
