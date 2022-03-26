package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{
			Name:    "Larah",
			City:    "Eusébio",
			Zipcode: "61760000",
		},
		{
			Name:    "Hallex",
			City:    "Eusébio",
			Zipcode: "61760000",
		},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
