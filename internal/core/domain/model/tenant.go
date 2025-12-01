package model

import (
	"time"

	"github.com/google/uuid"
)

// Tipo enumerado para países soportados
type Country string

const (
	CountryPeru     Country = "Perú"
	CountryEspana   Country = "España"
	CountryColombia Country = "Colombia"
	CountryChile    Country = "Chile"
)

type Tenant struct {
	ID           uuid.UUID
	Name         string
	Ruc          string
	EmailContact string
	PhoneContact string
	IsConfirm    bool
	IsActive     bool

	// Nuevos campos para dirección y branding
	Department string  // Departamento
	Province   string  // Provincia
	District   string  // Distrito
	Address    string  // Dirección detallada
	PostalCode string  // Código postal (opcional)
	LogoURL    string  // URL del logo (opcional)
	Country    Country // País

	CreatedAt time.Time
	UpdatedAt time.Time
}
