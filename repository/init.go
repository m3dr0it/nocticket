package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"noctiket/db"
	"noctiket/model/entity"
	"time"
)

var client *mongo.Client
var userCollection *mongo.Collection
var rolePermissionCollection *mongo.Collection
var ticketCollection *mongo.Collection
var sequenceCollection *mongo.Collection

func InitDbConnection(dbName string) {
	client = db.Connect()
	userCollection = client.Database(dbName).Collection("users")
	rolePermissionCollection = client.Database(dbName).Collection("role-permissions")
	sequenceCollection = client.Database(dbName).Collection("sequence")
	ticketCollection = client.Database(dbName).Collection("ticket")

	initSequences()
	addIndexes()
}

func initSequences() {
	for _, v := range entity.SEQUENCES {
		sequence, err := GetSequenceById(v)
		if sequence == 0 && err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				err := initializeSequenceById(v)
				if err != nil {
					log.Fatal(err.Error())
					return
				}
			}
		}
	}
}

func addIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	addUserIndex(ctx)
	addRolePermissionIndex(ctx)
}

func addRolePermissionIndex(ctx context.Context) {
	_, err := rolePermissionCollection.Indexes().DropAll(ctx)

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "role", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err = rolePermissionCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func addUserIndex(ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, // Field yang akan di-index
		Options: options.Index().SetUnique(true),
	}

	_, err := userCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal(err.Error())
	}
}
