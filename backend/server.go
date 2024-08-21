package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, world!")
	})

	// Start Server
	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}