package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getAddress() string {
	host, exist := os.LookupEnv("HOST")
	if !exist {
		host = "localhost"
	}
	port, exist := os.LookupEnv("PORT")
	if !exist {
		port = "4000"
	}
	address := fmt.Sprintf("%s:%s", host, port)
	return address
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	address := getAddress()
	log.Printf("Starting server on %s", address)
	err := http.ListenAndServe(address, mux)
	log.Fatal(err)
}
