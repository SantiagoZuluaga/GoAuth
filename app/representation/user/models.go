package user

type ResponseToken struct {
	Token string `json:"token"`
}

type ResponseError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
