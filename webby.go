package main

import (
    "log/slog"
    "net/http"
    "os"
    "gominiwiki/handlers"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5173"
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Wiki app starting", "time", time.Now().Format(time.RFC1123))

	mux := http.NewServeMux()
	mux.Handle("/style.css", http.FileServer(http.Dir("static")))
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	mux.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	mux.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))
	mux.HandleFunc("/create", handlers.GetNameHandler)
	mux.HandleFunc("/delete", handlers.DeleteHandler)

	logger.Info("Server ready", "url", "http://localhost:"+port)
	logger.Error("Server stopped", "err", http.ListenAndServe(":"+port, mux))
}

