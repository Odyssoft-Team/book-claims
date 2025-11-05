package http

import (
	"bytes"
	"claimbook-api/internal/infrastructure/jwt"
	"claimbook-api/pkg/util/apperror"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		var appErr *apperror.AppError

		if errors.As(err, &appErr) {
			logFn := logger.Info
			switch {
			case appErr.Code >= 500:
				logFn = logger.Error
			case appErr.Code >= 400:
				logFn = logger.Warn
			}

			fields := []zap.Field{
				zap.String("message", appErr.Message),
				zap.Int("status_code", appErr.Code),
				zap.String("file", appErr.File),
				zap.Int("line", appErr.Line),
			}

			if appErr.Underlying != nil {
				fields = append(fields, zap.NamedError("cause", appErr.Underlying))
			}
			if appErr.StackTrace != "" {
				fields = append(fields, zap.String("stacktrace", appErr.StackTrace))
			}
			fmt.Println("Middleware ejecutado con error:", appErr.Message)

			logFn("App error", fields...)
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
			return
		}

		logger.Error("Unhandled error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected internal error"})
	}
}

const maxLoggedBodySize = 2048

func RequestResponseLogger(logger *zap.Logger) gin.HandlerFunc {
	appEnv := os.Getenv("APP_ENV")

	logBodies := appEnv != "production"

	return func(c *gin.Context) {
		start := time.Now()

		var requestBody []byte
		if logBodies && c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				if len(bodyBytes) > maxLoggedBodySize {
					requestBody = append(bodyBytes[:maxLoggedBodySize], []byte("...[truncated]")...)
				} else {
					requestBody = bodyBytes
				}
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		respBody := &bytes.Buffer{}
		writer := &responseBodyWriter{body: respBody, ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		latency := time.Since(start)

		responseBody := respBody.Bytes()
		if len(responseBody) > maxLoggedBodySize {
			responseBody = append(responseBody[:maxLoggedBodySize], []byte("...[truncated]")...)
		}

		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.ByteString("request_body", requestBody),
			zap.ByteString("response_body", responseBody),
			zap.Duration("latency", latency),
		)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func AuthMiddleware(authLogger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			authLogger.Warn("Missing Authorization header",
				zap.String("ip", c.ClientIP()),
				zap.String("path", c.FullPath()),
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			authLogger.Warn("Invalid token format",
				zap.String("ip", c.ClientIP()),
				zap.String("auth_header", authHeader),
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		token := parts[1]
		claims, err := jwt.ValidateAccessToken(token)
		if err != nil {
			authLogger.Warn("Token validation failed",
				zap.String("ip", c.ClientIP()),
				zap.String("path", c.FullPath()),
				zap.Error(err),
			)
			var appErr *apperror.AppError
			if errors.As(err, &appErr) {
				c.AbortWithStatusJSON(appErr.Code, gin.H{"error": appErr.Message})
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			}
			return
		}
		fmt.Println("Middleware ejecutado con exito", claims.TenantID)

		if claims.RoleName == "" {
			authLogger.Warn("Token missing role_name claim", zap.String("ip", c.ClientIP()), zap.String("path", c.FullPath()))
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role information missing in token"})
			return
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, jwt.ContextKeyTenantID, claims.TenantID)
		ctx = context.WithValue(ctx, jwt.ContextKeyUserID, claims.UserID)
		ctx = context.WithValue(ctx, jwt.ContextKeyUserName, claims.UserName)
		ctx = context.WithValue(ctx, jwt.ContextKeyRoleName, claims.RoleName)
		c.Request = c.Request.WithContext(ctx)

		c.Set("user_id", claims.UserID)
		c.Set("tenant_id", claims.TenantID)
		c.Set("user_name", claims.UserName)
		c.Set("role_name", claims.RoleName)
		c.Set("claims", claims)

		authLogger.Info("Authentication successful",
			zap.String("user_id", fmt.Sprint(claims.UserID)),
			zap.String("user_name", claims.UserName),
			zap.String("role_name", claims.RoleName),
			zap.String("ip", c.ClientIP()),
			zap.String("path", c.FullPath()),
		)
		c.Next()
	}
}

func RoleAuthorizationMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleName, exists := c.Get("role_name")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role not found in context"})
			return
		}

		roleStr, ok := roleName.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid role format"})
			return
		}

		// Verifica si el role está entre los permitidos
		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied for role: " + roleStr})
	}
}
