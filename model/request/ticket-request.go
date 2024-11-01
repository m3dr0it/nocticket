package request

import "time"

type TicketRequest struct {
	TicketId       string        `form:"ticketId"`
	Title          string        `form:"title"`
	Description    string        `form:"description"`
	Status         string        `form:"status"`
	Priority       string        `form:"priority"`
	CreatedBy      string        `form:"createdBy"`
	AssignedTo     string        `form:"assignedTo"`
	ResolvedAtFrom time.Time     `form:"resolvedAtFrom"`
	ResolvedAtTo   time.Time     `form:"resolvedAtTo"`
	Resolution     string        `form:"resolution"`
	CreatedAtFrom  time.Time     `form:"createdAtFrom"`
	CreatedAtTo    time.Time     `form:"createdAtTo"`
	UpdatedAtFrom  time.Time     `form:"updatedAtFrom"`
	UpdatedAtTo    time.Time     `form:"updatedAtTo"`
	SLATimeBuffer  time.Duration `form:"slaTimeBuffer"`
}
