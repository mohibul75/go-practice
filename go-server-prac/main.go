package main

import (
	"fmt"
	"log"
	"net/http"
)

func  handleForm(w http.ResponseWriter, r *http.Request){

	 if err:= r.ParseForm(); err != nil{
		fmt.Fprintf(w, "parse error %v\n", err)
		return
	 }

	 fmt.Fprintf(w,"Post request successfully submitted\n")
	 name:= r.FormValue("name")
	 address:= r.FormValue("address")
	 fmt.Fprintf(w,"name = %s\n", name)
	 fmt.Fprintf(w, "address= %s\n", address)

}

func handleHello(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found\n", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not support", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w,"Hello WOrld")

}

func main ()  {

	fileServer:= http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", handleForm)

	fmt.Printf("Server Listen and Serve on port 8000")

	if err:= http.ListenAndServe(":8000",nil); err!= nil{
		log.Fatal(err)
	}
	
}