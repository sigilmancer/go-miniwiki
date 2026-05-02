package handlers

import (
    "net/http"
    "os"
    "fmt"
    "log"
)

// DeleteHandler removes a page file and redirects to home
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")
    if title == "" {
        http.Error(w, "Missing title", http.StatusBadRequest)
        return
    }

    filename := title + ".txt"
    if err := os.Remove(filename); err != nil {
        http.Error(w, fmt.Sprintf("Error deleting page %s: %v", title, err), http.StatusInternalServerError)
        return
    }

    log.Println("Deleted page:", title)
    http.Redirect(w, r, "/", http.StatusFound)
}