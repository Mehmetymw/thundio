package jwt

import "time"

type JWTConfig struct {
	SecretKey       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewJWTConfig(secretKey string, accessTokenTTL, refreshTokenTTL time.Duration) JWTConfig {
	return JWTConfig{
		SecretKey:       secretKey,
		AccessTokenTTL:  accessTokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
	}
}
