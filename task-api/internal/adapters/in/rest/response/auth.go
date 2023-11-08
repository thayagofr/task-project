package response

import (
	"task-api/internal/core/domain"
	"time"
)

type AccessToken struct {
	Token          string    `json:"token"`
	Type           string    `json:"type"`
	ExpirationDate time.Time `json:"duration"`
}

func FromAccessDomain(token *domain.AccessCredentials) *AccessToken {
	return &AccessToken{
		Token:          token.Token,
		Type:           token.Type,
		ExpirationDate: token.ExpirationDate,
	}
}
