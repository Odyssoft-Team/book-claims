param()

Write-Host "Generando Swagger docs..."

# Ejecuta swag usando go run
go run github.com/swaggo/swag/cmd/swag@v1.16.6 init -g ./cmd/app/main.go -o internal/infrastructure/http/docs

Write-Host "Swagger docs generados en internal/infrastructure/http/docs"