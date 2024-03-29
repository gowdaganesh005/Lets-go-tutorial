package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/gowdaganesh005/snippetbox/internals/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippets
	Snippets    []*models.Snippets
	Form any
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\pages\\*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles("C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\pages\\base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("C:\\Users\\gowda\\Desktop\\GO-project\\Lets-go-tutorial\\snippetbox\\ui\\html\\partials\\*.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
