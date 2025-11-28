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

type LocationHandler struct {
	locationUseCase *usecase.LocationUseCase
}

func NewLocationHandler(uc *usecase.LocationUseCase) *LocationHandler {
	return &LocationHandler{locationUseCase: uc}
}

// CreateLocation godoc
// @Summary Create a location for a tenant
// @Description Create location under a tenant
// @Tags location
// @Accept json
// @Produce json
// @Param id path string true "Tenant ID"
// @Param location body dto.CreateLocationDTO true "Location data"
// @Success 201 {object} model.Location
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tenant/{id}/location [post]
func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var createDto dto.CreateLocationDTO

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

	ctx := c.Request.Context()

	newLocation, err := h.locationUseCase.CreateLocation(ctx, &createDto)

	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, newLocation)
}

// GetLocationById godoc
// @Summary Get location by ID
// @Description Retrieve location by UUID
// @Tags location
// @Produce json
// @Param id path string true "Location ID"
// @Success 200 {object} model.Location
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /location/{id} [get]
func (h *LocationHandler) GetLocationById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	location, err := h.locationUseCase.GetLocationById(c.Request.Context(), id)
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, location)
}
