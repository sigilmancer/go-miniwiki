package handlers

import (
    "net/http"
    "gominiwiki/models"
)

// ViewHandler displays a page if it exists, or redirects to edit if not
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := models.LoadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}