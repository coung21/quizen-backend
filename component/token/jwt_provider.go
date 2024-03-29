package token

import (
	"fmt"
	"quizen/common"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenProvider interface {
	GenerateTokens(payload *TokenPayload) (*string, *string, error)
	Validate(myToken string) (*Claims, error)
	NewPayLoad(id uuid.UUID) *TokenPayload
}
type jwtProvider struct {
	secret        string
	accessExpiry  int
	refreshExpiry int
}

func NewJWTProvider(secret string, accessExpiry, refreshExpiry int) TokenProvider {
	return jwtProvider{
		secret:        secret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

type TokenPayload struct {
	ID uuid.UUID `json:"user_id"`
}

type Claims struct {
	jwt.RegisteredClaims
	ID uuid.UUID `json:"user_id"`
}

func (j jwtProvider) NewPayLoad(id uuid.UUID) *TokenPayload {
	return &TokenPayload{
		ID: id,
	}
}

func (j jwtProvider) GenerateTokens(payload *TokenPayload) (*string, *string, error) {
	now := time.Now()

	// Generate access token
	accessTokenClaims := Claims{
		ID: payload.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(j.accessExpiry))),
			ID:        fmt.Sprintf("%d", now.UnixNano()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(j.secret))
	if err != nil {
		return nil, nil, err
	}

	// Generate refresh token
	refreshTokenClaims := Claims{
		ID: payload.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(j.refreshExpiry))),
			ID:        fmt.Sprintf("%d", now.UnixNano()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(j.secret))
	if err != nil {
		return nil, nil, err
	}

	// Create token objects

	return &accessTokenString, &refreshTokenString, nil
}
func (j jwtProvider) Validate(myToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(myToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, common.InvalidJWTClaims
	}
	return claims, nil
}
