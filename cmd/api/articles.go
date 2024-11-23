package main

import (
	"fmt"
	"net/http"
)

func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new article")
}

func (app *application) showArticleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of article %d\n", id)
}
