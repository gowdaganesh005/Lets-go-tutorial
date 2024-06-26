package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gowdaganesh005/snippetbox/internals/models"
)

type application struct {
	errlog         *log.Logger
	infolog        *log.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
	user           *models.UserModel
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
	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		errlog:         errlog,
		infolog:        infoLog,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
		user:           &models.UserModel{DB: db},
	}
	tlsconfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:        *addr,
		ErrorLog:    errlog,
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		TLSConfig:   tlsconfig,
	}
	infoLog.Println("server running on ", *addr)

	err = srv.ListenAndServeTLS("C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\tls\\cert.pem", "C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\tls\\key.pem")
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
