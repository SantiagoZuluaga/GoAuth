package data

type User struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
