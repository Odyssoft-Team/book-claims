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

type TenantHandler struct {
	tenantUseCase *usecase.TenantUseCase
}

func NewTenantHandler(uc *usecase.TenantUseCase) *TenantHandler {
	return &TenantHandler{tenantUseCase: uc}
}

// CreateTenant godoc
// @Summary Create a new tenant
// @Description Create tenant with the provided data
// @Tags tenant
// @Accept json
// @Produce json
// @Param tenant body dto.CreateTenantDTO true "Tenant data"
// @Success 201 {object} model.Tenant
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tenant [post]
func (h *TenantHandler) CreateTenant(c *gin.Context) {
	var createDto dto.CreateTenantDTO

	if err := c.ShouldBindJSON(&createDto); err != nil {
		appErr := apperror.NewBadRequest("Invalid request body: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	ctx := c.Request.Context()

	newTenant, err := h.tenantUseCase.CreateTenant(ctx, &createDto)

	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, newTenant)
}

// GetAllTenants godoc
// @Summary List tenants
// @Description Get all tenants
// @Tags tenant
// @Produce json
// @Success 200 {array} model.Tenant
// @Failure 500 {object} map[string]string
// @Router /tenant [get]
func (h *TenantHandler) GetAllTenants(c *gin.Context) {
	// fallback: call usecase if exists
	// ...existing code...
}

// GetTenantById godoc
// @Summary Get tenant by ID
// @Description Get a tenant using its ID
// @Tags tenant
// @Produce json
// @Param id path string true "Tenant ID"
// @Success 200 {object} model.Tenant
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tenant/{id} [get]
func (h *TenantHandler) GetTenantById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	user, err := h.tenantUseCase.GetTenantById(c.Request.Context(), id)
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

// UpdateTenant godoc
// @Summary Update tenant
// @Description Update tenant fields (activate tenant etc.)
// @Tags tenant
// @Accept json
// @Produce json
// @Param id path string true "Tenant ID"
// @Param tenant body dto.UpdateTenantDTO true "Tenant update data"
// @Success 200 {object} model.Tenant
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tenant/{id} [patch]
func (h *TenantHandler) UpdateTenant(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	var updateDto dto.UpdateTenantDTO
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		appErr := apperror.NewBadRequest("Invalid request body: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	updatedComplaint, err := h.tenantUseCase.UpdateTenant(c.Request.Context(), id, &updateDto)
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, updatedComplaint)
}
