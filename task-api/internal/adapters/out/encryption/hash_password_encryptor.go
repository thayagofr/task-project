package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"task-api/internal/ports/out"
)

var _ out.PasswordEncryptor = &HashPasswordEncryptor{}

type HashPasswordEncryptor struct{}

func NewHashPasswordEncryptor() *HashPasswordEncryptor {
	return &HashPasswordEncryptor{}
}

func (h *HashPasswordEncryptor) Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *HashPasswordEncryptor) Compare(encrypted string, original string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(original)); err != nil {
		return false
	}
	return true
}
