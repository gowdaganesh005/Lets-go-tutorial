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

	router.Handler(http.MethodGet, "/", app.sessionManager.LoadAndSave(http.HandlerFunc(app.home)))
	router.Handler(http.MethodGet, "/snippet/view/:id", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetview)))
	router.Handler(http.MethodGet, "/snippet/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetCreate)))
	router.Handler(http.MethodPost, "/snippet/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetcreatePost)))
	router.Handler(http.MethodGet, "/user/signup", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userSignup)))
	router.Handler(http.MethodPost, "/user/signup", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userSignupPost)))
	router.Handler(http.MethodGet, "/user/login", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLogin)))
	router.Handler(http.MethodPost, "/user/login", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLoginPost)))
	router.Handler(http.MethodPost, "/user/logout", app.sessionManager.LoadAndSave(http.HandlerFunc(app.userLogout)))
	return app.recoverPanic(app.logRequest(secureHeaders(router)))

}
