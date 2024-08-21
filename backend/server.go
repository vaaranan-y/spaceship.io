package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Printf("Starting Server at Port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if(err != nil){
		log.Fatal(err)
	}
}