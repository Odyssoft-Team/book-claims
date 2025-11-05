package main

import (
	"claimbook-api/internal/config"
	"claimbook-api/internal/core/usecase"
	"claimbook-api/internal/infrastructure/http"
	"claimbook-api/internal/infrastructure/http/handler"
	"claimbook-api/internal/infrastructure/jwt"
	"claimbook-api/internal/infrastructure/logger"
	"claimbook-api/internal/infrastructure/persistence/database"
	apikey "claimbook-api/internal/infrastructure/persistence/postgres/api_key"
	"claimbook-api/internal/infrastructure/persistence/postgres/complaint"
	"claimbook-api/internal/infrastructure/persistence/postgres/location"
	"claimbook-api/internal/infrastructure/persistence/postgres/role"
	"claimbook-api/internal/infrastructure/persistence/postgres/session"
	"claimbook-api/internal/infrastructure/persistence/postgres/tenant"
	"claimbook-api/internal/infrastructure/persistence/postgres/user"

	"log"
	"os"

	"go.uber.org/zap"
)

func main() {
	cfg := config.LoadConfig()

	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	zapLogger := logger.NewZapLogger(cfg.Env, "logs/app.log")
	httpLogger := logger.NewZapLogger(cfg.Env, "logs/http.log")
	AuthLogger := logger.NewZapLogger(cfg.Env, "logs/auth.log")

	defer zapLogger.Sync()

	if err := jwt.InitKeys(); err != nil {
		zapLogger.Fatal("Error inicializando llaves JWT", zap.Error(err))
	}

	db, err := database.Connect(cfg.DB)
	if err != nil {
		zapLogger.Fatal("Database connection failed", zap.Error(err))
	}

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		if err := database.Migrate(db); err != nil {
			zapLogger.Fatal("Migration failed", zap.Error(err))
		}
		zapLogger.Info("Migrations completed")
	}

	complaintRepo := complaint.NewComplaintPGRepository(db)
	complaintUseCase := usecase.NewComplaintUseCase(complaintRepo)
	complaintHandler := handler.NewComplaintHandler(complaintUseCase)

	userRepo := user.NewUserPGRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	roleRepo := role.NewRolePGRepository(db)
	roleUseCase := usecase.NewRoleUseCase(roleRepo)
	roleHandler := handler.NewRoleHandler(roleUseCase)

	locationRepo := location.NewLocationPGRepository(db)
	locationUseCase := usecase.NewLocationUseCase(locationRepo)
	locationHandler := handler.NewLocationHandler(locationUseCase)

	sessionRepo := session.NewSessionPGRepository(db)
	authUseCase := usecase.NewAuthUseCase(sessionRepo, userRepo)
	sessionHandler := handler.NewAuthHandler(authUseCase)

	tenantRepo := tenant.NewTenantPGRepository(db)
	tenantUseCase := usecase.NewTenantUseCase(tenantRepo)
	tenantHandler := handler.NewTenantHandler(tenantUseCase)

	apiKeyRepo := apikey.NewApiKeyPGRepository(db)
	apiKeyUseCase := usecase.NewApiKeyUseCase(apiKeyRepo)
	apiKeyHandler := handler.NewApiKeyHandler(apiKeyUseCase)

	r := http.SetupRouter(complaintHandler, userHandler, roleHandler, locationHandler, sessionHandler, tenantHandler, apiKeyHandler, zapLogger, httpLogger, AuthLogger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
