package main

import (
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/login", logIn())
	http.HandleFunc("/signup", refresh())
	http.HandleFunc("/home",home())

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Error in server")
	}

}