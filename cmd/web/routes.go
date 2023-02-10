package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	// fileServer := http.FileServer(http.Dir("./ui/static/"))

	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/book/view", app.booksView)
	mux.HandleFunc("/book/create", app.bookCreate)
	mux.HandleFunc("/book/delete", app.bookDelete)
	mux.HandleFunc("/book/update", app.bookUpdate)
	return mux
}
