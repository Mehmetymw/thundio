package mqtt

import (
	"log"

	"github.com/mehmetymw/thundio/internal/auth/jwt"
)

type MQTTAuthHandler struct {
	tokenService *jwt.TokenService
}

func NewMQTTAuthHandler(tokenService *jwt.TokenService) *MQTTAuthHandler {
	return &MQTTAuthHandler{tokenService: tokenService}
}

func (h *MQTTAuthHandler) Authenticate(clientID, username, password string) bool {
	log.Printf("Authenticating client: %s", clientID)

	uniqueID, err := h.tokenService.ValidateToken(password)
	if err != nil {
		log.Printf("Authentication failed for client %s: %v", clientID, err)
		return false
	}

	if uniqueID != username {
		log.Printf("Invalid unique_id for client %s", clientID)
		return false
	}

	log.Printf("Authentication successful for client %s", clientID)
	return true
}
