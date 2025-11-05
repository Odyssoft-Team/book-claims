package usecase

import (
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/pkg/util/apperror"
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo port.UserRepository
}

func NewUserUseCase(repo port.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, userDTO *dto.CreateUserDTO) (*dto.UserResponseDTO, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperror.NewInternalError("failed to hash password", err)
	}
	userDTO.Password = string(hashedPassword)

	domainModel := mapper.CreateUserDTOToDomain(*userDTO)

	created, err := uc.userRepo.CreateUser(ctx, domainModel)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to create user", err)
	}
	resp := mapper.UserToResponseDTO(created)
	return &resp, nil
}

func (uc *UserUseCase) GetUserById(ctx context.Context, id uuid.UUID) (*dto.UserResponseDTO, error) {
	user, err := uc.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve user", err)
	}
	if user == nil {
		return nil, apperror.NewNotFoundError("User not found")
	}
	resp := mapper.UserToResponseDTO(user)
	return &resp, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, id uuid.UUID, updateDTO *dto.UpdateUserDTO) (*dto.UserResponseDTO, error) {
	user, err := uc.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve user for update", err)
	}
	if user == nil {
		return nil, apperror.NewNotFoundError("User not found")
	}

	mapper.UpdateUserFromDTO(user, *updateDTO)

	updated, err := uc.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to update user", err)
	}

	resp := mapper.UserToResponseDTO(updated)
	return &resp, nil
}

func (uc *UserUseCase) GetUsers(ctx context.Context) ([]dto.UserResponseDTO, error) {
	users, err := uc.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve users by tenant ID", err)
	}

	if len(users) == 0 {
		return nil, apperror.NewNotFoundError("No users found for the tenant")
	}

	var responses []dto.UserResponseDTO
	for _, user := range users {
		responses = append(responses, mapper.UserToResponseDTO(user))
	}

	return responses, nil
}

func (uc *UserUseCase) Login(context context.Context, username, password string) (*dto.UserResponseDTO, error) {
	user, err := uc.userRepo.FindByUserAuth(context, username)
	if err != nil {
		return nil, apperror.NewBadRequest("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, apperror.NewBadRequest("wrong password")
	}
	resp := mapper.UserToResponseDTO(user)
	return &resp, nil
}
