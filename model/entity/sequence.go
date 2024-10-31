package entity

const SEQ_TICKET_ID = "SEQ_TICKET_ID"

var SEQUENCES = []string{SEQ_TICKET_ID}

type Sequence struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	Sequence int64  `json:"sequence"`
}
