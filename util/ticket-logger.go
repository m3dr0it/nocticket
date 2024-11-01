package util

import (
	"fmt"
	"noctiket/model/entity"
	"time"
)

func LogTicketCreated(ticketId string, userEmail string) entity.LogEntry {
	return entity.LogEntry{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf("Ticket %s created", ticketId),
		UpdatedBy: userEmail,
	}
}

func LogUpdateStatus(ticketId string, userEmail string, prevStatus string, nextStatus string) entity.LogEntry {
	return entity.LogEntry{
		Timestamp: time.Now(),
		Message: fmt.Sprintf("Ticket %s update from %s to %s",
			ticketId, prevStatus, nextStatus),
		UpdatedBy: userEmail,
	}
}

func LogUpdateProgress(ticketId string, userEmail string, progressMessage string) entity.LogEntry {
	return entity.LogEntry{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf("Ticket %s progress : %s", ticketId, progressMessage),
		UpdatedBy: userEmail,
	}
}

func LogAssignTicket(ticketId string, userEmail string, assignee string) entity.LogEntry {
	return entity.LogEntry{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf("Ticket %s assigned to : %s", ticketId, assignee),
		UpdatedBy: userEmail,
	}
}

func LogCloseTicket(ticketId string, userEmail string, resolution string) entity.LogEntry {
	return entity.LogEntry{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf("Ticket %s is closed, resolution: %s", ticketId, resolution),
		UpdatedBy: userEmail,
	}
}
