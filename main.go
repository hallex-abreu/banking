package main

import (
	"fmt"
	"net/http"
)

func main() {

	//define routes
	http.HandleFunc("/greet", greet)

	//start server
	http.ListenAndServe("localhost:8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}
