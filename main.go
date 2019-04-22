package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Api service starting...")
	fmt.Println("Listening on port 8085")
	handleApiCall()
}

func handleApiCall() {

	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8085", nil)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "Welcome to HomePage\n")
	_, _ = fmt.Fprintf(w, "This app is to show the example of deployment")

}

