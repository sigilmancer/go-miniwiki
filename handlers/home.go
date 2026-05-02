package handlers

import (
    "net/http"
    "os"
    "strings"
)

// HomeHandler lists all saved pages and renders the home template
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    files, err := os.ReadDir(".")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var titles []string
    for _, file := range files {
        name := file.Name()
        if strings.HasSuffix(name, ".txt") {
            titles = append(titles, strings.TrimSuffix(name, ".txt"))
        }
    }

    // Render a template with the list of titles
    if err := templates.ExecuteTemplate(w, "home.html", titles); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}