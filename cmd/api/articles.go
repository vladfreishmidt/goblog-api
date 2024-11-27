package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vladfreishmidt/goblog-api/internal/data"
)

func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title        string            `json:"title"`
		PublishDate  time.Time         `json:"publish_date"`
		ReadTime     int32             `json:"read_time ,omitempty"`
		Authors      []string          `json:"authors"`
		Excerpt      string            `json:"excerpt"`
		Tags         []string          `json:"tags"`
		SourceURL    string            `json:"source_url"`
		RelatedLinks map[string]string `json:"related_links"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
		app.serverErrorResponse(w, r, err)
	}
}
