package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gowdaganesh005/snippetbox/internals/models"
)

type application struct {
	errlog        *log.Logger
	infolog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "http network")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := OpenDB(*dsn)
	if err != nil {
		errlog.Fatal(err)
	}
	db.Ping()

	defer db.Close()
	templateCache, err := newTemplateCache()
	if err != nil {
		errlog.Fatal(err)
	}

	app := &application{
		errlog:   errlog,
		infolog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache, 
	}

	infoLog.Println("server running on ", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errlog,
		Handler:  app.routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		errlog.Println("error running the server ", err)

	}

}
func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
