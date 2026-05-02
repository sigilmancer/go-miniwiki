package handlers

import (
    "html/template"
    "net/http"
    "regexp"
    "gominiwiki/utils"
    "gominiwiki/models"
)

var templates = template.Must(template.ParseFiles(
    "templates/edit.html",
    "templates/view.html",
    "templates/home.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9\\-]+)$")

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        slug := utils.Slugify(m[2])
        fn(w, r, slug)
    }
}