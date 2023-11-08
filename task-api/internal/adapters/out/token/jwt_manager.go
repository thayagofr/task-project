package token

import (
	"github.com/dgrijalva/jwt-go"
	"task-api/internal/core/domain"
	"task-api/internal/ports/out"
	"time"
)

const (
	tokenType = "Bearer"
)

var _ out.TokenManager = &JWTManager{}

type JWTManager struct {
	secret string
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{secret: secret}
}

func (provider *JWTManager) Generate(user domain.User) (*domain.AccessCredentials, error) {
	var (
		expirationTime = time.Now().Add(time.Hour * 1)
		claims         = jwt.MapClaims{
			"sub": user.ID,
			"exp": expirationTime,
			"iat": time.Now().Unix(),
			"user": map[string]interface{}{
				"email":           user.Email,
				"registered_date": user.RegisteredDate,
			},
		}
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(provider.secret))
	if err != nil {
		return nil, out.ErrGeneratingToken
	}

	return &domain.AccessCredentials{
		Token:          signedToken,
		Type:           tokenType,
		ExpirationDate: expirationTime,
	}, nil
}

func (provider *JWTManager) Validate(token string) error {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, out.ErrInvalidTokenSigningMethod
		}
		return []byte(provider.secret), nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return out.ErrInvalidToken
	}

	return nil
}
