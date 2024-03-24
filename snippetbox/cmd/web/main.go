package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)
type application struct{
	errlog *log.Logger
	infolog *log.Logger

}

func main() {
	addr := flag.String("addr", ":4000", "http network")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	

	app:=&application{
		errlog:errlog,
		infolog:infoLog,
	}
	
	
	infoLog.Println("server running on ", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errlog,
		Handler:  app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		errlog.Println("error running the server ", err)

	}

}
