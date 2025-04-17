package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/pages", handler)
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {
    target := r.URL.Query().Get("target")
    if target == "" {
        http.Error(w, "Missing target URL", http.StatusBadRequest)
        return
    }

    pages := Crawl(target)
    fmt.Fprintf(w, "%v\n", pages)
}

