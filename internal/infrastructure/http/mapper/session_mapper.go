package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"

	"github.com/google/uuid"
)

func CreateSessionDTOToDomain(c dto.CreateSessionDTO) *model.Session {
	return &model.Session{
		ID:           uuid.New(),
		UserID:       c.UserID,
		TenantID:     c.TenantID,
		RefreshToken: c.RefreshToken,
		IP:           c.IP,
		UserAgent:    c.UserAgent,
		ExpiresAt:    c.ExpiresAt,
		Revoked:      false,
	}
}

func SessionToResponseDTO(session *model.Session) *dto.ResponseSessionDTO {
	return &dto.ResponseSessionDTO{
		ID:           session.ID,
		UserID:       session.UserID,
		TenantID:     session.TenantID,
		RefreshToken: session.RefreshToken,
		IP:           session.IP,
		UserAgent:    session.UserAgent,
		CreatedAt:    session.CreatedAt,
		UpdatedAt:    session.UpdatedAt,
		ExpiresAt:    session.ExpiresAt,
		Revoked:      session.Revoked,
	}
}

func UpdateSessionFromDTO(existing *model.Session, dto dto.UpdateSessionDTO) {
	if dto.Revoked != nil {
		existing.Revoked = *dto.Revoked
	}
}
