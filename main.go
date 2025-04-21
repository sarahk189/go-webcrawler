package main

import (
	"fmt"
	"net/http"
)


func main() {
    http.HandleFunc("/pages", CrawlHandler)
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}



