package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Use the notFound() helper
		return
	}
	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err) // Use the serverError() helper.
	// 	return
	// }
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err) // Use the serverError() helper.
	// }
	fmt.Println("Endpoint Hit: home")
}
func (app *application) booksView(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}
func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		return
	}
	fmt.Println("Endpoint Hit: create book")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	// update our global Articles array to include
	// our new Article
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}
