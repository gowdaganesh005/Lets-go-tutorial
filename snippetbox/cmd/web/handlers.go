package main

import (
	"fmt"

	"html/template"

	"net/http"
	"strconv"
)

func (app application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notfound(w)
		return
	}
	files := []string{
		"C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\pages\\base.html",
		"C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\pages\\home.html",
		"C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\partials\\nav.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {

		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Write([]byte("hello hi this is a snippet box"))
}



func (app application) snippetcreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title:="0 snail"
   	content:="0 snail\nClimb mount fuji,\nBut slowly\nBut slowly,slowly!\n \n -Kobayassi Issa"
  	expires:=7
	id,err:=app.snippets.Insert(title,content,expires)
	if err!=nil {
		app.serverError(w,err)
		return


	}
	http.Redirect(w,r,fmt.Sprintf("/snippet/view?id=%d",id),http.StatusSeeOther)
	
	w.Write([]byte("create the snippet"))
}
func (app application) snippetview(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notfound(w)
		return

	}
	fmt.Fprint(w, "snippet view for id ", id)
}
