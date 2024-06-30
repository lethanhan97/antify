package main

import (
	"antify/cmd"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := cmd.NewRouter()

	// TODO: Make port an ENV Variable
	fmt.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
