package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\static\\"))

	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", home)
	router.HandleFunc("/snippet/view", snippetview)
	router.HandleFunc("/snippet/create", snippetcreate)
	fmt.Print("server running on 4000")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Fatal("error running the server")

	}

}
