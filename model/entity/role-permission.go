package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type RolePermission struct {
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Role string             `json:"role" bson:"role"`
	Apis []Api              `json:"apis" bson:"apis"`
}

type Api struct {
	Method string `json:"method" bson:"method"`
	Path   string `json:"path" bson:"path"`
}
