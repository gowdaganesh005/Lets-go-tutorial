package main

import "net/http"

func (app application) routes() *http.ServeMux {
	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\static\\"))

	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", app.home)
	router.HandleFunc("/snippet/view", app.snippetview)
	router.HandleFunc("/snippet/create", app.snippetcreate)
	return router


}
