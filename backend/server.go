package main

import (
	"fmt"
	"log"
	"net/http"
)

func testEndpoint(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path != "/hello"){
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if(r.Method != "GET"){
		http.Error(w, "Unsupported Method", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, world!")
}

func main(){
	http.HandleFunc("/hello", testEndpoint)

	// Start Server
	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}