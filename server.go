package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", SimpleServer)
	http.ListenAndServe(":8080", nil)
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ole Out Please %s", r.URL.Path[1:])
}
