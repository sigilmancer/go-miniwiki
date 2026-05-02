package handlers

import (
    "net/http"
    "gominiwiki/utils"

)

func GetNameHandler(w http.ResponseWriter, r *http.Request) {
    rawTitle := r.FormValue("title")
    if rawTitle == "" {
        http.Error(w, "Title is required", http.StatusBadRequest)
        return
    }

    safeTitle := utils.Slugify(rawTitle) // use utils package
    if safeTitle == "" {
        http.Error(w, "Invalid title", http.StatusBadRequest)
        return
    }
	
    http.Redirect(w, r, "/edit/"+safeTitle, http.StatusFound)
}