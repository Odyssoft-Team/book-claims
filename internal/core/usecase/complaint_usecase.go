package usecase

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/pkg/util/apperror"
	"context"
	"time"

	"github.com/google/uuid"
)

type ComplaintUseCase struct {
	complaintRepo         port.ComplaintRepository
	complaintSequenceRepo port.ComplaintSequenceRepository
}

func NewComplaintUseCase(repo port.ComplaintRepository, sequence port.ComplaintSequenceRepository) *ComplaintUseCase {
	return &ComplaintUseCase{
		complaintRepo:         repo,
		complaintSequenceRepo: sequence,
	}
}

func (uc *ComplaintUseCase) CreateComplaint(ctx context.Context, complaintDTO *dto.CreateComplaintDTO) (*dto.ComplaintResponse, error) {

	codePublic, err := uc.complaintSequenceRepo.GenerateCodePublic(ctx, complaintDTO.TenantID, "EMPRESA")
	if err != nil {
		return nil, apperror.NewInternalError("cannot generate complaint code", err)
	}
	complaintDTO.CodePublic = codePublic

	domainModel := mapper.CreateComplaintDTOToDomain(*complaintDTO)

	created, err := uc.complaintRepo.CreateComplaint(ctx, domainModel)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to create complaint", err)
	}
	resp := mapper.ComplaintToResponseDTO(created)
	return &resp, nil
}

func (uc *ComplaintUseCase) GetComplaintByCodePublic(ctx context.Context, code string) (*dto.ComplaintResponse, error) {
	complaint, err := uc.complaintRepo.GetByPublicCode(ctx, code)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve complaint", err)
	}
	if complaint == nil {
		return nil, apperror.NewNotFoundError("Complaint not found")
	}
	resp := mapper.ComplaintToResponseDTO(complaint)
	return &resp, nil
}

func (uc *ComplaintUseCase) GetComplaintById(ctx context.Context, id uuid.UUID) (*dto.ComplaintResponse, error) {
	complaint, err := uc.complaintRepo.GetComplaintById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve complaint", err)
	}
	if complaint == nil {
		return nil, apperror.NewNotFoundError("Complaint not found")
	}
	resp := mapper.ComplaintToResponseDTO(complaint)
	return &resp, nil
}

func (uc *ComplaintUseCase) UpdateComplaint(ctx context.Context, id uuid.UUID, updateDTO *dto.UpdateComplaintDTO) (*dto.ComplaintResponse, error) {
	complaint, err := uc.complaintRepo.GetComplaintById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve complaint for update", err)
	}
	if complaint == nil {
		return nil, apperror.NewNotFoundError("Complaint not found")
	}

	// Aplicar cambios desde DTO
	mapper.UpdateComplaintFromDTO(complaint, *updateDTO)

	// Si la respuesta fue enviada, fijar metadata
	if updateDTO.ResponseStatus != nil && *updateDTO.ResponseStatus == string(model.RESPONSE_SENT) {
		now := time.Now().UTC()
		complaint.ResponseSentAt = &now
		// si no se proporcionó ResponderID en DTO, intentar obtener de contexto (no disponible aquí)
		// cambiar estado a ATENDIDO si está en RECIBIDO/EVALUACION/PROCESO
		if complaint.Status == model.RECIBIDO || complaint.Status == model.EVALUACION || complaint.Status == model.PROCESO {
			complaint.Status = model.ATENDIDO
		}
	}

	// new_status ya fue aplicado en mapper

	updated, err := uc.complaintRepo.UpdateComplaint(ctx, complaint)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to update complaint", err)
	}

	resp := mapper.ComplaintToResponseDTO(updated)
	return &resp, nil
}

func (uc *ComplaintUseCase) GetComplaints(ctx context.Context) ([]dto.ComplaintResponse, error) {
	complaints, err := uc.complaintRepo.GetComplaints(ctx)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve complaints by tenant ID", err)
	}

	if len(complaints) == 0 {
		return nil, apperror.NewNotFoundError("No complaints found for the tenant")
	}

	var responses []dto.ComplaintResponse
	for _, complaint := range complaints {
		responses = append(responses, mapper.ComplaintToResponseDTO(complaint))
	}

	return responses, nil
}

func (uc *ComplaintUseCase) GetSummaryReport(ctx context.Context) (model.SummaryReport, error) {
	return uc.complaintRepo.GetSummary(ctx)
}
