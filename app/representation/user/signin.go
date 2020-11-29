package user

import (
	"encoding/json"
	"net/http"

	"github.com/SantiagoZuluaga/GoAuth/app/domain"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	user, message, status := domain.IsUserValid(
		"signin",
		r,
	)
	if message != "" {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}

	jwt, err := domain.GenerateJWT(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseToken{
		Token: jwt,
	})
}

func SignInView(w http.ResponseWriter, r *http.Request) {

}
