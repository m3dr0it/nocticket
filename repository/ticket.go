package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"noctiket/model/entity"
	"noctiket/model/request"
	"time"
)

func GenerateTicketId() (string, error) {
	const prefix = "T"
	date := time.Now().Format("20060102")

	seq, err := GetNexSequenceById(entity.SEQ_TICKET_ID)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	seqStr := fmt.Sprintf("%05d", seq)
	id := prefix + date + seqStr
	return id, nil
}

func SaveTicket(ticket entity.Ticket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ticketCollection.InsertOne(ctx, ticket)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GetTickets(ticketRequest request.TicketRequest) ([]entity.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := ticketCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var tickets []entity.Ticket
	err = cursor.All(ctx, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func GetTicketById(id string) (entity.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var ticket entity.Ticket

	err := ticketCollection.FindOne(ctx, bson.D{
		{"ticket_id", id},
	}, nil).Decode(&ticket)

	if err != nil {
		return entity.Ticket{}, err
	}

	return ticket, nil
}

func UpdateTicket(ticket entity.Ticket) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define the update fields
	update := bson.M{
		"$set": bson.M{
			"assigned_to": ticket.AssignedTo,
			"status":      ticket.Status,
			"updated_at":  time.Now(),
		},
	}

	updatedRes, err := ticketCollection.UpdateOne(ctx,
		bson.M{"ticket_id": ticket.TicketId},
		update)

	if err != nil {
		log.Fatal(err)
		return err
	}

	if updatedRes.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
