package entity

import "time"

type Ticket struct {
	TicketId    string     `json:"ticketId" bson:"ticket_id"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Status      string     `json:"status" bson:"status"`
	Priority    string     `json:"priority" bson:"priority"`
	CreatedAt   time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" bson:"updated_at"`
	CreatedBy   string     `json:"createdBy" bson:"created_by"`
	AssignedTo  string     `json:"assignedTo" bson:"assigned_to"`
	ResolvedAt  time.Time  `json:"resolvedAt" bson:"resolved_at"`
	Resolution  string     `json:"resolution" bson:"resolution"`
	SLATime     time.Time  `json:"slaTime" bson:"sla_time"`
	Log         []LogEntry `json:"log" bson:"log"`
}

type LogEntry struct {
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`  // Time of log entry
	Message   string    `json:"message" bson:"message"`      // Log message
	UpdatedBy string    `json:"updatedBy" bson:"updated_by"` // Who made the change
}
