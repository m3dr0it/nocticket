package request

type AssignTicket struct {
	TicketId string `json:"ticket_id"`
	Assignee string `json:"assignee"`
}
