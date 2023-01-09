package main

import (
	"fmt"
	"net/http"
)

func getSystemStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("System Specs")
}

func registerClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Client")
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Client")
}

func returnClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Return Client")
}

func initServer() {
	http.HandleFunc("/system", getSystemStatus)
	http.HandleFunc("/registerClient", registerClient)
	http.HandleFunc("/deleteClient", deleteClient)
	http.HandleFunc("/returnClient", returnClient)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP Error: %v\n", err)
	}
}
