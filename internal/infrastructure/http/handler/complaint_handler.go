package handler

import (
	"claimbook-api/internal/core/usecase"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/pkg/util/apperror"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ComplaintHandler struct {
	complaintUseCase *usecase.ComplaintUseCase
}

func NewComplaintHandler(uc *usecase.ComplaintUseCase) *ComplaintHandler {
	return &ComplaintHandler{complaintUseCase: uc}
}

func (h *ComplaintHandler) CreateComplaint(c *gin.Context) {
	var createDto dto.CreateComplaintDTO

	if err := c.ShouldBindJSON(&createDto); err != nil {
		appErr := apperror.NewBadRequest("Invalid request body: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	if apiKeyID, exists := c.Get("api_key_id"); exists {
		if id, ok := apiKeyID.(uuid.UUID); ok {
			createDto.ApiKeyID = id
			fmt.Println("API Key ID:", id.String())

		}
	}

	ctx := c.Request.Context()

	newComplaint, err := h.complaintUseCase.CreateComplaint(ctx, &createDto)

	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, newComplaint)
}

func (h *ComplaintHandler) GetComplaintByCodePublic(c *gin.Context) {
	code := c.Param("code")

	complaint, err := h.complaintUseCase.GetComplaintByCodePublic(c.Request.Context(), code)
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, complaint)
}

func (h *ComplaintHandler) GetComplaintById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	complaint, err := h.complaintUseCase.GetComplaintById(c.Request.Context(), id)
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, complaint)
}

func (h *ComplaintHandler) UpdateComplaint(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		appErr := apperror.NewBadRequest("Invalid UUID format: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	var updateDto dto.UpdateComplaintDTO
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		appErr := apperror.NewBadRequest("Invalid request body: " + err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	updatedComplaint, err := h.complaintUseCase.UpdateComplaint(c.Request.Context(), id, &updateDto)
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

func (h *ComplaintHandler) GetComplaints(c *gin.Context) {
	complaints, err := h.complaintUseCase.GetComplaints(c.Request.Context())
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, complaints)
}

func (h *ComplaintHandler) SummaryReportHandler(c *gin.Context) {
	ctx := c.Request.Context()

	summary, err := h.complaintUseCase.GetSummaryReport(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get summary"})
		return
	}

	c.JSON(http.StatusOK, summary)
}
