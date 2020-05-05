package models

// User describes the user domain. 
// Ideally we don't want the password to be passed around the system, even though it is hidden from JSON. 
// An alternative would be to define multiple User structs, so the Password field would be removed when listing user for example.  
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
