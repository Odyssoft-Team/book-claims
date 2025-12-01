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
		Type:            model.ComplaintType(c.Type),
		Status:          model.ComplaintStatus(c.Status),
		CategoryID:      c.CategoryID,
		Source:          model.ComplaintSource(c.Source),
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

	var responseSentAt string
	if complaint.ResponseSentAt != nil && !complaint.ResponseSentAt.IsZero() {
		responseSentAt = complaint.ResponseSentAt.UTC().Format(time.RFC3339)
	} else {
		responseSentAt = ""
	}

	var responderIDStr string
	if complaint.ResponderID != nil {
		responderIDStr = complaint.ResponderID.String()
	} else {
		responderIDStr = ""
	}

	var updatedAtStr string
	if complaint.UpdatedAt != nil {
		updatedAtStr = complaint.UpdatedAt.UTC().Format(time.RFC3339)
	} else {
		updatedAtStr = ""
	}

	return dto.ComplaintResponse{
		ID:              complaint.ID,
		TenantID:        complaint.TenantID,
		LocationID:      complaint.LocationID,
		Type:            string(complaint.Type),
		Status:          string(complaint.Status),
		CategoryID:      complaint.CategoryID,
		Source:          string(complaint.Source),
		ApiKeyID:        complaint.ApiKeyID,
		CodePublic:      complaint.CodePublic,
		Description:     complaint.Description,
		RequestedAction: complaint.RequestedAction,
		ResponseText:    complaint.ResponseText,
		ResponseStatus:  string(complaint.ResponseStatus),
		ResponderID:     responderIDStr,
		ResponseSentAt:  responseSentAt,
		CreatedAt:       complaint.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:       updatedAtStr,
		ResolvedAt:      resolvedAtStr,
		IsClosed:        complaint.IsClosed,
	}
}

func UpdateComplaintFromDTO(existing *model.Complaint, upd dto.UpdateComplaintDTO) {
	// Cambiar status por new_status
	if upd.NewStatus != nil {
		existing.Status = model.ComplaintStatus(*upd.NewStatus)
	}

	// Actualizar respuesta: draft o send
	if upd.ResponseText != nil {
		existing.ResponseText = *upd.ResponseText
	}
	if upd.ResponseStatus != nil {
		existing.ResponseStatus = model.ResponseStatus(*upd.ResponseStatus)
		// si es SENT, se debe fijar ResponseSentAt (se hará en usecase con timestamp actual)
	}
	if upd.ResponderID != nil {
		existing.ResponderID = upd.ResponderID
	}
}
