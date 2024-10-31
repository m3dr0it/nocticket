package request

import "time"

type UserRequest struct {
	Id            string    `json:"id" form:"id"`
	Email         string    `json:"email" form:"email"`
	Password      string    `json:"password" form:"password"`
	Role          string    `json:"role" form:"role"`
	IsActive      bool      `json:"isActive" form:"isActive"`
	CreatedAtFrom time.Time `json:"createdAtFrom" form:"createdAtFrom"`
	CreatedAtTo   time.Time `json:"createdAtTo" form:"createdAtTo"`
}
