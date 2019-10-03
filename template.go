package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type TemplateHandler struct {
	once     sync.Once
	filename string
	html     *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		path := filepath.Join("views", t.filename)
		t.html = template.Must(template.ParseFiles(path))
	})
	t.html.Execute(w, nil)
}

func render(view string) *TemplateHandler {
	return &TemplateHandler{filename: view}
}
