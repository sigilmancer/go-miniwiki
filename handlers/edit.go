package handlers

import (
    "net/http"
    "gominiwiki/models"
)

// EditHandler loads a page for editing, or creates a new one if not found
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := models.LoadPage(title)
    if err != nil {
        p = &models.Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}