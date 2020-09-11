package models

type Customer struct {
	ID             int    `json:"id,omitempty"`
	FName          string `json:"first_name" form:"first_name"`
	LName          string `json:"last_name" form:"last_name"`
	Email          string `json:"email" form:"email"`
	Mobile         string `json:"mobile" form:"mobile"`
	HashedPassword string
}
