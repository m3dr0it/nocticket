package request

type TicketProgressRequest struct {
	TicketId string `json:"ticketId"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}
