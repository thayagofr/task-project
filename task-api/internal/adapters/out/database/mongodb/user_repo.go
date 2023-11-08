package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"task-api/internal/adapters/out/database/mongodb/documents"
	"task-api/internal/core/domain"
	"task-api/internal/ports/out"
)

var _ out.UserProvider = &UserMongoDBProvider{}

type UserMongoDBProvider struct {
	collection *mongo.Collection
}

func (repo *UserMongoDBProvider) FindCredentialsByEmail(ctx context.Context, email string) (*domain.UserCredentials, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *UserMongoDBProvider) RegisterUser(ctx context.Context, newUser *domain.UserToRegister) (*domain.RegisteredUser, error) {
	userDoc := documents.FromDomain(newUser)
	result, err := repo.collection.InsertOne(ctx, userDoc)
	if err != nil {
		return nil, err
	}
	return &domain.RegisteredUser{
		ID:             result.InsertedID.(string),
		Email:          newUser.Email,
		RegisteredDate: userDoc.CreationDate,
	}, nil
}

func (repo *UserMongoDBProvider) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	err := repo.collection.FindOne(ctx, bson.M{
		"notification": email,
	}).Err()

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
