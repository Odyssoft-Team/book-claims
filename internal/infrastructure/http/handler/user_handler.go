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

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: uc}
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user for a tenant
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDTO true "User data"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var createDto dto.CreateUserDTO

	if err := c.ShouldBindJSON(&createDto); err != nil {
		appErr := apperror.NewBadRequest("Invalid request body: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	ctx := c.Request.Context()

	newUser, err := h.userUseCase.CreateUser(ctx, &createDto)

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

// GetUserById godoc
// @Summary Get user by ID
// @Description Retrieve user by UUID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [get]
func (h *UserHandler) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	user, err := h.userUseCase.GetUserById(c.Request.Context(), id)
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

// Login godoc
// @Summary User login
// @Description Authenticate user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.AuthRequest true "Credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /user/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var login dto.AuthRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.Error(apperror.NewBadRequest("Invalid JSON: " + err.Error()))
		return
	}

	data, err := h.userUseCase.Login(c.Request.Context(), login.UserName, login.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, data)
}
