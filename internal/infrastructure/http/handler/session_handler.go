package handler

import (
	"claimbook-api/internal/core/usecase"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/pkg/util/apperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(uc *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: uc}
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var request dto.AuthRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "detail": err.Error()})
		return
	}

	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()

	response, err := h.authUseCase.Login(c.Request.Context(), request, ip, userAgent)
	if err != nil {
		if appErr, ok := err.(*apperror.AppError); ok {
			c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	// Refresh token (long-lived, e.g., 7 days)
	refreshCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    response.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   60 * 60 * 24 * 7, // 7 days
	}

	c.Writer.Header().Add("Set-Cookie", refreshCookie.String())

	// Devolver la respuesta
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "Login successful",
		"access_token": response.AccessToken,
	})
}

func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	// 1. Intentar leer el refresh token
	cookie, err := c.Request.Cookie("refresh_token")
	if err == nil {
		// Invalidate it on the backend only if it exists
		_ = h.authUseCase.Logout(c.Request.Context(), cookie.Value)
	}

	// 2. Delete refresh_token cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   -1, // expires immediately
	})

	// 3. Confirm logout
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout successful",
	})
}
