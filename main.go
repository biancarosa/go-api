package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func ready(w http.ResponseWriter, r *http.Request) {
	// check db connection and other componentes
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/health-check", healthCheck)
	http.HandleFunc("/ready-check", ready)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Listening on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
