package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"noctiket/model/entity"
	"time"
)

func AddRolePermission(permission entity.RolePermission) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rolePermissionCollection.InsertOne(ctx, bson.M{
		"role": permission.Role,
		"apis": permission.Apis,
	})

	if err != nil {
		return err
	}

	return nil
}

func GetRolePermission(role string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}

	if role != "" {
		filter["role"] = role
	}

	find, err := rolePermissionCollection.Find(ctx, filter)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var res []bson.M
	err = find.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetRolePermissionById(id string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panicln(err.Error())
		return nil, err
	}
	filter = append(filter, bson.E{Key: "_id", Value: objID})

	find, err := rolePermissionCollection.Find(ctx, filter)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var res []bson.M
	err = find.All(ctx, &res)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return res, nil
}

func DeleteRolePermission(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Panicln(err.Error())
		return err
	}

	filter = append(filter, bson.E{Key: "_id", Value: objID})
	_, err = rolePermissionCollection.DeleteOne(ctx, filter)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
