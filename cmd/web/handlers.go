package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	reqBody, _ := io.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	// update our global Articles array to include
	// our new Article
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

func (app *application) bookDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id != "" {
		for index, book := range Books {
			if book.ISBN == id {
				Books = append(Books[:index], Books[index+1:]...)
			}
		}
	}

}

func (app *application) bookUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id != "" {
		var updatedEvent Book
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &updatedEvent)
		for i, book := range Books {
			if book.ISBN == id {
				book.Title = updatedEvent.Title
				book.Description = updatedEvent.Description
				book.Author = updatedEvent.Author
				book.Published = updatedEvent.Published
				Books[i] = book
				json.NewEncoder(w).Encode(book)
			}
		}
	}
}
