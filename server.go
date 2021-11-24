package main

import (
	"net/http"

	"github.com/abmutungi/go-server/test"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello!")
// }

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST request successful")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")

// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Address = %s\n", address)
// }

// func main() {
// 	fileServer := http.FileServer(http.Dir("./static")) // creates the file server object using the FileServer function.

// 	http.Handle("/", fileServer)                        // the Handle route, which accepts a path and the fileserver

// 	http.HandleFunc("/hello", helloHandler) // Update this line of code

// 	http.HandleFunc("/form", formHandler)

// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
func main() {
	srv := test.NewServer()
	http.ListenAndServe(":8080", srv)
}
