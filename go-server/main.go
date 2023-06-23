package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {

	if err:= r.ParseForm(); err!= nil {
		fmt.Fprintf(w, "ParseForm error %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successfully submitted\n")
	name:= r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w, " the name of %s\n", name)
	fmt.Fprintf(w, " the address of %s\n", address)
}

func handleHello(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "hello world")

}

func main(){

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", handleForm)


	fmt.Printf("Server Listen and Serve at 8000")
	if err := http.ListenAndServe(":8000",nil); err!=nil {
		log.Fatal(err)
	}


}