package main

import (
	"fmt"
	"log"
	"net/http"
)

func heolloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello......!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", heolloHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8443", nil); err != nil {
		log.Fatal(err)
	}

}
