package main

import (
	"fmt"
	"log"
	"net/http"
)

// Function to handle hello end point

func heolloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
	fmt.Println("Entered hello handler")
	fmt.Fprintf(w, "Hello......!")
}

// Main function
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", heolloHandler)
	fmt.Println("Starting server at port 8443")
	// You can change listening port here
	fmt.Println("Listening on port 8443")
	if err := http.ListenAndServe(":8443", nil); err != nil {
		log.Fatal(err)
	}

}
