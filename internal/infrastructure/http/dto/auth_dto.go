package dto

// AuthRequest representa las credenciales para autenticación
// swagger:model AuthRequest
type AuthRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse representa los tokens de autenticación
// swagger:model AuthResponse
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
