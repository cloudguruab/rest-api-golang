package main

import (
	"fmt"
	"log"
	"net/http"
)


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the home page, hello world.")
	fmt.Println("/homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}