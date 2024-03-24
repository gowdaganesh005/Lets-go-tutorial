package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errlog.Println(trace)
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	app.errlog.Output(2,trace)
}

func (app *application) clientError(w http.ResponseWriter,status int){
	http.Error(w,http.StatusText(status),status)


}

func (app *application) notfound(w http.ResponseWriter){
	app.clientError(w,http.StatusNotFound)
}
