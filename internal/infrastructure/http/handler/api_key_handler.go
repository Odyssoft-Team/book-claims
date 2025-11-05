package handler

import (
	"claimbook-api/internal/core/usecase"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/pkg/util/apperror"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApiKeyHandler struct {
	apiKeyUseCase *usecase.ApiKeyUseCase
}

func NewApiKeyHandler(uc *usecase.ApiKeyUseCase) *ApiKeyHandler {
	return &ApiKeyHandler{apiKeyUseCase: uc}
}

func (h *ApiKeyHandler) CreateApiKey(c *gin.Context) {
	var createDto dto.CreateApiKeyDTO

	if err := c.ShouldBindJSON(&createDto); err != nil {
		appErr := apperror.NewBadRequest("Invalid request body: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	tenantIDStr := c.Param("id")
	tenantID, err := uuid.Parse(tenantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	createDto.TenantID = tenantID
	createDto.ApiKey = uuid.NewString()

	ctx := c.Request.Context()

	newUser, err := h.apiKeyUseCase.CreateApiKey(ctx, &createDto)

	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (h *ApiKeyHandler) GetApiKeyById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	user, err := h.apiKeyUseCase.GetApiKeyById(c.Request.Context(), id)
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
