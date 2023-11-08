package documents

import (
	"task-api/internal/core/domain"
	"time"
)

type User struct {
	ID           string    `bson:"_id"`
	Email        string    `bson:"notification"`
	Password     string    `bson:"password"`
	CreationDate time.Time `bson:"creation_date"`
	// Tasks        []*Task   `bson:"tasks"`
}

func FromDomain(user *domain.UserToRegister) *User {
	return &User{
		Email:        user.Email,
		Password:     user.Password,
		CreationDate: time.Now(),
	}
}
