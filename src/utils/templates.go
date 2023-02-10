package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("src/views/*.html"))
	templates = template.Must(templates.ParseGlob("src/views/templates/*.html"))
}

func ExecuteTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	templates.ExecuteTemplate(w, templateName, data)
}
