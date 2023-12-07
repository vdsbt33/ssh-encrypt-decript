package main

import (
    //"crypto"
    "errors"
    "fmt"
    "net/http"
    "os"
)


func serveIndex(w http.ResponseWriter, r *http.Request) {
    allowedFiles := []string { "/index.html", "/style.css", "/script.js" }

    if len(r.URL.Path) <= 1 {
        http.ServeFile(w, r, "./src/static/index.html")
        return
    }

    for _, filePath := range allowedFiles {
        if r.URL.Path == filePath {
            http.ServeFile(w, r, "./src/static" + filePath)
            return
        }
    }
}

func serveHttp() {
    fmt.Println("Serving html in port 8080")

    http.HandleFunc("/", serveIndex)
    err := http.ListenAndServe(":8080", nil)

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("Server closed") 
    } else if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
        os.Exit(1)
    } 
    
}
