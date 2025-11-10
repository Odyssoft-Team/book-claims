package http

import (
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter(
	complaintHandler *handler.ComplaintHandler,
	userHandler *handler.UserHandler,
	roleHandler *handler.RoleHandler,
	locationHandler *handler.LocationHandler,
	sessionHandler *handler.AuthHandler,
	tenantHandler *handler.TenantHandler,
	apiKeyHandler *handler.ApiKeyHandler,
	apiKeyRepo port.ApiKeyRepository,
	logger *zap.Logger,
	httpLogger *zap.Logger,
	authLogger *zap.Logger,
) *gin.Engine {
	router := gin.Default()
	router.Use(RequestResponseLogger(httpLogger))
	router.Use(ErrorLoggerMiddleware(logger))

	publicApi := router.Group("/api/v1")
	privateApi := router.Group("/api/v1")
	privateApi.Use(AuthMiddleware(authLogger))

	{
		publicComplaint := publicApi.Group("/complaint")
		publicComplaint.Use(ApiKeyMiddleware(apiKeyRepo, httpLogger))
		{
			publicComplaint.POST("", complaintHandler.CreateComplaint)
			publicComplaint.GET("/code/:code", complaintHandler.GetComplaintByCodePublic)

		}
		privateComplaint := privateApi.Group("/complaint")
		{
			privateComplaint.GET("", complaintHandler.GetComplaints)
			privateComplaint.GET("/:id", complaintHandler.GetComplaintById)
			privateComplaint.POST("/:id/action", complaintHandler.UpdateComplaint)

		}
	}
	{
		report := privateApi.Group("/report")
		{
			report.GET("/summary", complaintHandler.SummaryReportHandler)
		}
	}
	{
		publicUser := publicApi.Group("/user")
		{
			publicUser.POST("/login", userHandler.Login)
			publicUser.POST("/", userHandler.CreateUser)
		}

		privateUser := privateApi.Group("/user")
		{
			privateUser.GET("/:id", userHandler.GetUserById)
		}
	}

	{
		publicRole := publicApi.Group("/role")
		{
			publicRole.POST("/", roleHandler.CreateRole)
		}
		privateRole := privateApi.Group("/role")
		{
			privateRole.GET("/:id", roleHandler.GetRoleById)
		}
	}
	{
		location := privateApi.Group("/location")
		{
			location.POST("/", locationHandler.CreateLocation)
			location.GET("/:id", locationHandler.GetLocationById)
		}
	}
	{
		auth := publicApi.Group("/auth")
		{
			auth.POST("/login", sessionHandler.LoginHandler)
			auth.POST("/logout", sessionHandler.LogoutHandler)
		}

	}
	{
		publicTenant := publicApi.Group("/tenant")
		{
			publicTenant.POST("/", tenantHandler.CreateTenant)
			publicTenant.GET("/:id", tenantHandler.GetTenantById)
			publicTenant.PATCH("/:id", tenantHandler.UpdateTenant)
			publicTenant.POST("/:id/location", locationHandler.CreateLocation)
			publicTenant.POST("/:id/api-keys", apiKeyHandler.CreateApiKey)
		}
	}
	{
		apiKey := privateApi.Group("api_key")
		{
			apiKey.POST("/", apiKeyHandler.CreateApiKey)
			apiKey.GET("/:id", apiKeyHandler.GetApiKeyById)
		}
	}
	return router
}
