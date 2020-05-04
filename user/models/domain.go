package models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
