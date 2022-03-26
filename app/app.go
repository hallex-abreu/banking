package app

import (
	"log"
	"net/http"

	"github.com/hallex-abreu/banking/app/handlers"
)

func Start() {

	mux := http.NewServeMux()

	//define routes
	mux.HandleFunc("/greet", handlers.Greet)
	mux.HandleFunc("/customers", handlers.GetAllCustomers)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
