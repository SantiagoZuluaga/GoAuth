package domain

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/SantiagoZuluaga/GoAuth/app/data"
)

func IsEmailValid(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 || len(email) > 254 {
		return false
	}

	return emailRegex.MatchString(email)
}

func IsUserValid(endpoint string, r *http.Request) (data.User, string, int) {

	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

	}

	isEmailValid := IsEmailValid(user.Email)
	if !isEmailValid {
		return data.User{}, "INVALID EMAIL", 403
	}

	if user.Password == "" {
		return data.User{}, "INVALID PASSWORD", 403
	}

	if endpoint == "signup" {
		if user.Username == "" {
			return data.User{}, "INVALID USERNAME", 403
		}

		if user.Password != user.ConfirmPassword {
			return data.User{}, "PASSWORD DON'T MACH", 403
		}
	}

	return user, "", http.StatusOK

}
