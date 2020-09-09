package models

type Customer struct {
	ID       int    `json:"id,omitempty"`
	FName    string `json:"first_name"`
	LName    string `json:"last_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
