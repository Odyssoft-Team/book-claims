#!/usr/bin/env bash
set -euo pipefail

echo "Generando Swagger docs..."

# Ejecuta swag usando go run para evitar depender de instalación global
go run github.com/swaggo/swag/cmd/swag@v1.16.6 init -g ./cmd/app/main.go -o internal/infrastructure/http/docs

echo "Swagger docs generados en internal/infrastructure/http/docs"