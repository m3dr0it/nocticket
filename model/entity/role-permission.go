package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

type RoleGroup []Role

type RolePermission struct {
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Role Role               `json:"role" bson:"role"`
	Apis []Api              `json:"apis" bson:"apis"`
}

type Api struct {
	Method string `json:"method" bson:"method"`
	Path   string `json:"path" bson:"path"`
}

func (rg RoleGroup) Contains(role Role) bool {
	for _, r := range rg {
		if r == role {
			return true
		}
	}
	return false
}
