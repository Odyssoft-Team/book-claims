package usecase

import (
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/internal/infrastructure/jwt"
	"claimbook-api/pkg/util/apperror"
	"context"
)

type AuthUseCase struct {
	sessionRepo port.SessionRepository
	userRepo    port.UserRepository
}

func NewAuthUseCase(repo port.SessionRepository, userRepo port.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		sessionRepo: repo,
		userRepo:    userRepo}
}

func (uc *AuthUseCase) Login(ctx context.Context, request dto.AuthRequest, ip, userAgent string) (*dto.AuthResponse, error) {
	user, err := uc.userRepo.FindByUserAuth(ctx, request.UserName)
	if err != nil {
		return nil, err
	}

	if user.Email == "" || user.ID.String() == "" {
		return nil, apperror.NewInternalError("user does not have a valid username or ID", nil)
	}
	var fullName = user.FullName

	// 1. Generate access and refresh tokens.
	accessToken, err := jwt.GenerateAccessToken(user.Email, user.ID.String(), fullName, user.TenantID.String(), user.RoleID.String(), user.LocationID.String(), user.RoleName)
	if err != nil {
		return nil, apperror.NewInternalError("error generating access token", err)
	}

	refreshToken, refreshTokenExp, err := jwt.GenerateRefreshToken(user.Email, user.TenantID.String())
	if err != nil {
		return nil, apperror.NewInternalError("error generating refresh token", err)
	}

	sessionDTO := dto.CreateSessionDTO{
		UserID:       user.ID,
		TenantID:     user.TenantID,
		RefreshToken: refreshToken,
		IP:           ip,
		UserAgent:    userAgent,
		ExpiresAt:    refreshTokenExp,
	}

	session := mapper.CreateSessionDTOToDomain(sessionDTO)

	_, err = uc.sessionRepo.Create(ctx, session)
	if err != nil {
		return nil, apperror.NewInternalError("error registering session", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *AuthUseCase) RefreshToken(ctx context.Context, refreshToken string) (*dto.AuthResponse, error) {
	claims, err := jwt.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, apperror.NewBadRequest("Invalid refresh token")
	}

	email, ok := claims["sub"].(string)
	if !ok || email == "" {
		return nil, apperror.NewBadRequest("Refresh token without valid 'subject'")
	}

	tenant, ok := claims["tenantId"].(string)
	if !ok || tenant == "" {
		return nil, apperror.NewBadRequest("Refresh token without valid 'tenant'")
	}

	session, err := uc.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, apperror.NewNotFoundError("Session not found or expired")
	}

	if session.Revoked {
		return nil, apperror.NewBadRequest("Session has been revoked")
	}

	userInfo, err := uc.userRepo.GetUserById(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	var fullName = userInfo.FullName

	userId := session.UserID.String()
	locationId := userInfo.LocationID.String()
	newAccessToken, err := jwt.GenerateAccessToken(email, userId, fullName, tenant, userInfo.RoleID.String(), locationId, userInfo.RoleName)
	if err != nil {
		return nil, apperror.NewInternalError("Error generating new access token", err)
	}

	return &dto.AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *AuthUseCase) Logout(ctx context.Context, refreshToken string) error {
	_, err := jwt.ValidateRefreshToken(refreshToken)
	if err != nil {
		return apperror.NewBadRequest("invalid refresh token")
	}

	session, err := uc.sessionRepo.FindByRefreshToken(ctx, refreshToken)
	if err != nil {
		return apperror.NewNotFoundError("session not found or already expired")
	}

	if session.Revoked {
		return apperror.NewBadRequest("the session has already been revoked")
	}

	session.Revoked = true
	if err := uc.sessionRepo.Update(ctx, session); err != nil {
		return apperror.NewInternalError("error revoking session", err)
	}

	return nil
}
