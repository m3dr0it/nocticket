package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
