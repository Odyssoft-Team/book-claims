package jwt

import (
	"claimbook-api/pkg/util/apperror"
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

const ContextUserClaimsKey = "userClaims"

type CustomClaims struct {
	UserID   string    `json:"user_id"`
	UserName string    `json:"name"`
	TenantID uuid.UUID `json:"tenant_id"`
	RoleName string    `json:"role_name"`
	jwt.RegisteredClaims
}

type contextKey string

// Declarar tus claves con tipo seguro
const (
	ContextKeyTenantID contextKey = "tenant_id"
	ContextKeyUserID   contextKey = "user_id"
	ContextKeyUserName contextKey = "user_name"
	ContextKeyRoleName contextKey = "role_name"
)

var accessExpiration = time.Duration(3) * time.Minute
var refreshExpiration = time.Duration(7*24) * time.Hour

var keysPath = os.Getenv("JWT_KEYS_PATH")

func InitKeys() error {
	if keysPath == "" {
		keysPath = "keys" // default
	}

	privBytes, err := os.ReadFile(fmt.Sprintf("%s/private.pem", keysPath))
	if err != nil {
		return fmt.Errorf("error leyendo private.pem: %w", err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return fmt.Errorf("error parseando private.pem: %w", err)
	}

	pubBytes, err := os.ReadFile(fmt.Sprintf("%s/public.pem", keysPath))
	if err != nil {
		return fmt.Errorf("error leyendo public.pem: %w", err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return fmt.Errorf("error parseando public.pem: %w", err)
	}

	return nil
}

func ValidateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token inválido: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("no se pudieron extraer los claims del token")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return nil, fmt.Errorf("token expirado")
		}
	} else {
		return nil, fmt.Errorf("el token no tiene un campo de expiración válido")
	}

	return claims, nil
}

func ValidateAccessToken(tokenString string) (*CustomClaims, error) {
	if publicKey == nil {
		return nil, apperror.NewInternalError("Public key not initialized", nil)
	}

	// Parse el token usando CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, apperror.NewUnauthorized(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, apperror.NewUnauthorized(fmt.Sprintf("Invalid token: %v", err))
	}

	// Validar que sea un token correcto y que los claims sean del tipo esperado
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, apperror.NewUnauthorized("Invalid or malformed token")
	}

	return claims, nil
}

func GenerateAccessToken(username string, userId string, name string, tenantId string, roleId string, locationId string, roleName string) (string, error) {
	claims := jwt.MapClaims{
		"sub":         username,
		"name":        name,
		"tenant_id":   tenantId,
		"location_id": locationId,
		"role_id":     roleId,
		"role_name":   roleName,
		"user_id":     userId,
		"tokenType":   "access",
		"exp":         time.Now().Add(accessExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func GenerateRefreshToken(username string, tenantId string) (string, time.Time, error) {
	expirationTime := time.Now().Add(refreshExpiration)

	claims := jwt.MapClaims{
		"sub":       username,
		"tenantId":  tenantId,
		"tokenType": "refresh",
		"exp":       expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error al firmar el refresh token: %w", err)
	}

	return signedToken, expirationTime, nil
}
