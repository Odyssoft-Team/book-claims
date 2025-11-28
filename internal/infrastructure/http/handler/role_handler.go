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

type RoleHandler struct {
	roleUseCase *usecase.RoleUseCase
}

func NewRoleHandler(uc *usecase.RoleUseCase) *RoleHandler {
	return &RoleHandler{roleUseCase: uc}
}

// CreateRole godoc
// @Summary Create a role for a tenant
// @Description Create role
// @Tags role
// @Accept json
// @Produce json
// @Param role body dto.CreateRoleDTO true "Role data"
// @Success 201 {object} model.Role
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /role [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var createDto dto.CreateRoleDTO

	if err := c.ShouldBindJSON(&createDto); err != nil {
		appErr := apperror.NewBadRequest("Invalid request body: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	ctx := c.Request.Context()

	newRole, err := h.roleUseCase.CreateRole(ctx, &createDto)

	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, newRole)
}

// GetRoleById godoc
// @Summary Get role by ID
// @Description Retrieve role by UUID
// @Tags role
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} model.Role
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /role/{id} [get]
func (h *RoleHandler) GetRoleById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	role, err := h.roleUseCase.GetRoleById(c.Request.Context(), id)
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, role)
}
