package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/home",handler);
	http.HandleFunc("/some",handler2);
	log.Fatal(http.ListenAndServe("localhost:8888",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Response for path /home");
}


func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Response for path /some");
}