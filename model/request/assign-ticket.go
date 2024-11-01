package request

type AssignTicket struct {
	TicketId string `json:"ticketId"`
	Assignee string `json:"assignee"`
}
