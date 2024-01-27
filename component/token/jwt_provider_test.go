package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenProvider(t *testing.T) {
	TokenProvider := NewJWTProvider("secret", 1, 1)

	t.Run("GenerateTokens", func(t *testing.T) {
		payload := TokenPayload{
			ID: 1,
		}

		accessToken, refreshToken, err := TokenProvider.GenerateTokens(&payload)

		assert.NoError(t, err, "GenerateTokens() error")
		assert.NotNil(t, accessToken, "GenerateTokens() accessToken")
		assert.NotNil(t, refreshToken, "GenerateTokens() refreshToken")
	})

	t.Run("NewPayLoad", func(t *testing.T) {
		payload := TokenProvider.NewPayLoad(1)

		assert.NotNil(t, payload, "NewPayLoad() payload")
	})

	t.Run("Validate", func(t *testing.T) {
		payload := TokenPayload{
			ID: 1,
		}

		accessToken, _, err := TokenProvider.GenerateTokens(&payload)
		assert.NoError(t, err, "GenerateTokens() error")

		claims, err := TokenProvider.Validate(accessToken.Token)
		assert.NoError(t, err, "Validate() error")
		assert.NotNil(t, claims, "Validate() claims")
	})

}
