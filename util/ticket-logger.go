package util

import (
	"fmt"
	"log"
	"noctiket/constant"
	"noctiket/model/entity"
	"time"
)

func LogTicket(ticket entity.Ticket, user entity.User, logStatus string) entity.LogEntry {
	var message string
	switch logStatus {
	case constant.Open:
		message = fmt.Sprintf("Ticket %s created", ticket.TicketId)
		break
	case constant.OnProgress:
		message = fmt.Sprintf("Ticket %s is updated from %s to %s",
			ticket.TicketId,
			ticket.Status,
			logStatus)
		log.Println(message)
		break
	case constant.Close:
		message = fmt.Sprintf("Ticket %s is closed", ticket.TicketId)
		break
	case constant.AssignTo:
		message = fmt.Sprintf("Ticket %s is assigned to %s", ticket.TicketId, ticket.AssignedTo)
		break

	default:
		message = fmt.Sprintf("Ticket %s is %s", logStatus)
	}

	return entity.LogEntry{
		Timestamp: time.Now(),
		Message:   message,
		UpdatedBy: user.Email,
	}
}
