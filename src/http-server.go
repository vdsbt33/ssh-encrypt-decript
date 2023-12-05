package main

import (
    //"crypto"
    "errors"
    "fmt"
    "net/http"
    "os"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./src/static/index.html")
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
