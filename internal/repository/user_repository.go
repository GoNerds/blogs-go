package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"blogs-go/internal/models"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)

	return err
}

func (r *UserRepository) GetUsers() (users []models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)
	if err != nil {
		return
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return &user, err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.D{
			{Key: "name", Value: user.Name},
			{Key: "email", Value: user.Email},
		},
	})

	return err
}

func (r *UserRepository) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}