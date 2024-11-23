package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	config JWTConfig
}

func NewTokenService(config JWTConfig) *TokenService {
	return &TokenService{config: config}
}

func (t *TokenService) GenerateAccessToken(uniqueID string) (string, error) {
	claims := jwt.MapClaims{
		"unique_id": uniqueID,
		"exp":       time.Now().Add(t.config.AccessTokenTTL).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.config.SecretKey))
}

func (t *TokenService) GenerateRefreshToken(uniqueID string) (string, error) {
	claims := jwt.MapClaims{
		"unique_id": uniqueID,
		"exp":       time.Now().Add(t.config.RefreshTokenTTL).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.config.SecretKey))
}

func (t *TokenService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.config.SecretKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", jwt.ErrTokenMalformed
	}

	uniqueID, ok := claims["unique_id"].(string)
	if !ok {
		return "", jwt.ErrTokenInvalidClaims
	}

	return uniqueID, nil
}

func ExtractTokenFromPayload(payload string) string {
	parts := strings.Split(payload, ":")
	if len(parts) == 2 && strings.TrimSpace(parts[0]) == "token" {
		return strings.TrimSpace(parts[1])
	}
	return ""
}
