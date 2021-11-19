package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", SimpleServer)
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ole Out Please %s", r.URL.Path[1:])
}
