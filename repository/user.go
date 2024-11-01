package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"noctiket/model/entity"
	"noctiket/model/request"
	"noctiket/util"
	"time"
)

func AddUser(user entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, bson.M{
		"email":     user.Email,
		"password":  util.MD5Hash(user.Password),
		"role":      user.Role,
		"isActive":  user.IsActive,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	})
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	filter := bson.D{
		{"email", email},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, filter, nil).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUsers(request request.UserRequest) ([]bson.M, error) {
	projection := bson.D{
		{"password", 0},
	}
	filter := bson.D{}

	if request.Id != "" {
		objID, err := primitive.ObjectIDFromHex(request.Id)
		if err != nil {
			log.Panicln(err.Error())
			return nil, err
		}
		filter = append(filter, bson.E{Key: "_id", Value: objID})
	}

	if request.Email != "" {
		filter = append(filter, bson.E{Key: "email", Value: request.Email})
	}

	if request.Role != "" {
		filter = append(filter, bson.E{Key: "role", Value: request.Role})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	find, err := userCollection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		log.Panicln(err.Error())
		return nil, err
	}

	var result []bson.M
	if err = find.All(ctx, &result); err != nil {
		log.Panicln(err.Error())
		return nil, err
	}

	return result, nil
}
