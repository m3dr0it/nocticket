package request

import "time"

type TicketRequest struct {
	TicketId      string    `json:"ticketId"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	Priority      string    `json:"priority"`
	CreatedBy     string    `json:"createdBy"`
	AssignedTo    string    `json:"assignedTo"`
	ResolvedAt    time.Time `json:"resolvedAt"`
	Resolution    string    `json:"resolution"`
	CreatedAtTo   time.Time `json:"createdAtTo"`
	UpdatedAtFrom time.Time `json:"updatedAtFrom"`
	UpdatedAtTo   time.Time `json:"updatedAtTo"`
	SLATimeFrom   time.Time `json:"slaTime"`
}
