package main

import (
	"fmt"

	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\pages\\base.html",
		"C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\pages\\home.html",
		"C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\partials\\nav.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal("internal server error :", err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, fmt.Sprint("internalserver error: ", err), http.StatusBadGateway)
	}
	w.Write([]byte("hello hi this is a snippet box"))
}
func snippetcreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}

	w.Write([]byte("create the snippet"))
}
func snippetview(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return

	}
	fmt.Fprint(w, "snippet view for id ", id)
}
