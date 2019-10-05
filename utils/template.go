package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type ContextFunc = func(r *http.Request) interface{}

type TemplateHandler struct {
	once     sync.Once
	filename string
	context  ContextFunc
	html     *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	t.once.Do(func() {
		path := filepath.Join("views", t.filename)
		t.html = template.Must(template.ParseFiles(path))
	})
	if t.context != nil {
		data = t.context(r)
	}
	t.html.Execute(w, data)
}

func Render(view string, context ContextFunc) *TemplateHandler {
	return &TemplateHandler{filename: view, context: context}
}
