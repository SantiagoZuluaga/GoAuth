package representation

import (
	"encoding/json"
	"net/http"

	"github.com/SantiagoZuluaga/GoAuth/app/data"
	"github.com/SantiagoZuluaga/GoAuth/app/domain"
	"github.com/gorilla/mux"
)

type ResponseToken struct {
	Token string `json:"token"`
}

func indexHandler(response http.ResponseWriter, request *http.Request) {

}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error")
		return
	}

	jwt, err := domain.GenerateJWT(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error")
		return
	}

	jwt, err := domain.GenerateJWT(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

func validateHandler(w http.ResponseWriter, r *http.Request) {
	valid, message := domain.ValidateToken(r)
	if !valid {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func Router() *mux.Router {

	router := mux.NewRouter()
	//router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/auth/signin", signinHandler).Methods("POST")
	router.HandleFunc("/api/auth/signup", signupHandler).Methods("POST")
	router.HandleFunc("/api/auth/validate", validateHandler).Methods("GET")
	return router
}
