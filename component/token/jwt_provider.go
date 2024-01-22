package token

import (
	"fmt"
	"quizen/common"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenProvider interface {
	GenerateToken(payload *TokenPayload, expiry int) (*Token, error)
	Validate(myToken string) (*Claims, error)
	NewPayLoad(id int, role string) *TokenPayload
}
type jwtProvider struct {
	secret string
}

type TokenPayload struct {
	ID int `json:"user_id"`
}

type Claims struct {
	jwt.RegisteredClaims
	ID int `json:"user_id"`
}
type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"token_created"`
	Expiry  int       `json:"token_expiry"`
}

func (j *jwtProvider) NewPayLoad(id int, role string) *TokenPayload {
	return &TokenPayload{
		ID: id,
	}
}

func (j *jwtProvider) GenerateToken(payload *TokenPayload, expiry int) (*Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID: payload.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * time.Duration(expiry))),

			ID: fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	aToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:   aToken,
		Created: now,
		Expiry:  expiry,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*Claims, error) {
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
