package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
	// "bookstore.hoomans.dev/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

type Book struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	Description string    `json:"desc"`
	Published   time.Time `json:"pub"`
}

var Books []Book

func main() {

	Books = []Book{
		{Title: "Off the clock", Author: "Laura Vanderkam", ISBN: "978-03494-2233-6", Description: "Hello", Published: time.Now()},
	}
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
