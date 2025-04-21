package main

import (
    "encoding/json"
    "net/http"
)

// CrawlHandler handles GET /pages?target=...
func CrawlHandler(w http.ResponseWriter, r *http.Request) {
    target := r.URL.Query().Get("target")
    if target == "" {
        http.Error(w, "Missing target URL", http.StatusBadRequest)
        return
    }

    pages := Crawl(target)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(struct {
        Domain string   `json:"domain"`
        Pages  []string `json:"pages"`
    }{
        Domain: target,
        Pages:  pages,
    })
}
