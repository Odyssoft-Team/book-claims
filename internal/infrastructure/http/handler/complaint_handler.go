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

// CreateComplaint godoc
// @Summary Create a complaint
// @Description Create a complaint (public endpoint using API Key)
// @Tags complaint
// @Accept json
// @Produce json
// @Param complaint body dto.CreateComplaintDTO true "Complaint data"
// @Success 201 {object} model.Complaint
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /complaint [post]
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

// GetComplaintByCodePublic godoc
// @Summary Get complaint by public code
// @Description Retrieve a complaint using public code (API Key required)
// @Tags complaint
// @Produce json
// @Param code path string true "Public code"
// @Success 200 {object} model.Complaint
// @Failure 404 {object} map[string]string
// @Router /complaint/code/{code} [get]
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

// GetComplaintById godoc
// @Summary Get complaint by ID
// @Description Retrieve complaint by UUID
// @Tags complaint
// @Produce json
// @Param id path string true "Complaint ID"
// @Success 200 {object} model.Complaint
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /complaint/{id} [get]
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

// UpdateComplaint godoc
// @Summary Update complaint
// @Description Update complaint status or details
// @Tags complaint
// @Accept json
// @Produce json
// @Param id path string true "Complaint ID"
// @Param complaint body dto.UpdateComplaintDTO true "Update data"
// @Success 200 {object} model.Complaint
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /complaint/{id}/action [post]
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

// GetComplaints godoc
// @Summary List complaints
// @Description List all complaints (requires auth)
// @Tags complaint
// @Produce json
// @Success 200 {array} model.Complaint
// @Failure 500 {object} map[string]string
// @Router /complaint [get]
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

// SummaryReportHandler godoc
// @Summary Get summary report
// @Description Returns summary report data
// @Tags report
// @Produce json
// @Success 200 {object} model.SummaryReport
// @Failure 500 {object} map[string]string
// @Router /report/summary [get]
func (h *ComplaintHandler) SummaryReportHandler(c *gin.Context) {
	ctx := c.Request.Context()

	summary, err := h.complaintUseCase.GetSummaryReport(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get summary"})
		return
	}

	c.JSON(http.StatusOK, summary)
}
