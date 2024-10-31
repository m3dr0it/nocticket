package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"noctiket/model/entity"
	"time"
)

func GetSequenceById(id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{
		"id": id,
	}

	var sequence entity.Sequence
	err := sequenceCollection.FindOne(ctx, filter, nil).Decode(&sequence)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return sequence.Sequence, nil
}

func GetNexSequenceById(id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"id": id,
	}

	update := bson.M{
		"$inc": bson.M{
			"sequence": 1,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var sequence entity.Sequence
	err := sequenceCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&sequence)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return sequence.Sequence, nil
}

func initializeSequenceById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := sequenceCollection.InsertOne(ctx, bson.M{
		"id":       id,
		"sequence": 0,
	}, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
