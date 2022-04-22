package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	server := &http.Server{
		Addr:     getAddress(),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", server.Addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
