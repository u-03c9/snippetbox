package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/u-03c9/snippetbox/pkg/models/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
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

func getDsn() string {
	dbString, exist := os.LookupEnv("DSN")
	if !exist {
		dbString = "web:pass@/snippetbox?parseTime=true"
	}
	return dbString
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(getDsn())
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	server := &http.Server{
		Addr:     getAddress(),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", server.Addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
