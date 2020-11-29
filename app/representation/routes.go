package representation

import (
	"encoding/json"
	"net/http"

	"github.com/SantiagoZuluaga/GoAuth/app/domain"
	"github.com/SantiagoZuluaga/GoAuth/app/representation/user"
	"github.com/gorilla/mux"
)

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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ERROR 404")
}

func CORS(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		next(w, r)
	})
}

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/auth/signin", CORS(user.SignInHandler)).Methods("POST")
	router.HandleFunc("/api/auth/signup", CORS(user.SignUpHandler)).Methods("POST")
	router.HandleFunc("/api/auth/validate", validateHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	return router
}
