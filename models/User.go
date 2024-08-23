package models

type User struct {
	ID        string `json:"id" validate:"required,uuid4"`
	FirstName string `json:"first_name" validate:"required,min=2,max=20"`
	LastName  string `json:"last_name" validate:"required,min=2,max=20"`
	Biography string `json:"biography" validate:"required,min=2,max=450"`
}
