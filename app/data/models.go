package data

type User struct {
	Username        string `json:"username,omitempty"`
	Email           string `json:"email"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmpassword,omitempty"`
}
