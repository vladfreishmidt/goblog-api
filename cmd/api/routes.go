package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/articles", app.createArticleHandler)
	router.HandlerFunc(http.MethodGet, "/v1/articles/:id", app.showArticleHandler)

	return router
}
