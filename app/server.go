package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SantiagoZuluaga/GoAuth/app/representation"
)

func RunServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	router := representation.Router()
	fmt.Println("Server running in 	port: http://localhost:" + port)
	http.ListenAndServe(port, router)
}
