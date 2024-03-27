package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
    cfg := getConfig()
    flag.Parse()
    mux := http.NewServeMux()

    fileServer := http.FileServer(neuteredFileSystem{http.Dir(cfg.staticDir)})
    mux.Handle("/static", http.NotFoundHandler())
    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)
    mux.HandleFunc("POST /snippet/create", snippetCreatePost)

    log.Printf("starting server on %s\n", cfg.addr)
    
    err := http.ListenAndServe(cfg.addr, mux)
    log.Fatal(err)
}

