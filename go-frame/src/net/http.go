package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get!")
}

func set(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "set!")
}
