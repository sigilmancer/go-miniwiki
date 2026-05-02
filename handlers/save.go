package handlers

import (
    "net/http"
    "log"             
    "gominiwiki/models"
)

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &models.Page{Title: title, Body: []byte(body)} // use models.Page
    err := p.Save() // capitalized method
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("Saving page with title:", title)
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

