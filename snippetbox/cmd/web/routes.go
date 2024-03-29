package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\static\\"))

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notfound(w)
	})
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetview)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetcreatePost)
	return app.recoverPanic(app.logRequest(secureHeaders(router)))

}
