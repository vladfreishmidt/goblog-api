package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vladfreishmidt/goblog-api/internal/data"
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

	article := data.Article{
		ID:          id,
		Title:       "Go Turns 15",
		PublishDate: time.Date(2024, time.Month(10), 15, 0, 0, 0, 0, time.UTC),
		SourceURL:   "https://go.dev/blog/15years",
		ReadTime:    1,
		Tags:        []string{"functions", "anniversary"},
		Authors:     []string{"Austin Clements"},
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"article": article}, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
