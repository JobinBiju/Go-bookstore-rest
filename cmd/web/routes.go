package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/book/view", app.booksView)
	mux.HandleFunc("/book/create", app.bookCreate)
	mux.HandleFunc("/book/update", app.bookUpdate)
	mux.HandleFunc("/book/delete", app.bookDelete)
	return mux
}
