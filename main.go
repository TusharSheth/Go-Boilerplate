package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	port := ":8080"
	fmt.Println("Server is up and running at", port)
	log.Fatal(http.ListenAndServe(port, router))
}
