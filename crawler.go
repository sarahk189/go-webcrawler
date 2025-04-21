package main

// //These are tools we're borrowing from Go's toolbox.
// Each one adds special abilities to our program.
import (
   "fmt"                     // Lets us print messages (like a log)
    "net/http"               // Lets us make web requests (go visit a page)
    "golang.org/x/net/html"  // Lets us read and understand HTML (like reading page code)
    "net/url"                // Helps us work with URLs (fixing links, comparing domains)
)

//map to keep track of visited pages
var visited = make(map[string]bool)

func Crawl(startUrl string) []string {
    visited = make(map[string]bool) // reset the map for each crawl, starting fresh
    base, _ := url.Parse(startUrl) // lets you acess and manipulate components of the URL
    crawl(startUrl, base)
    
    pages := []string{} // slice to store the visited pages
    for page := range visited {  //iterating over the map to get the visited pages
        pages = append(pages, page) //adding the visited pages to a slice
    }
    return pages //returning the slice of visited pages
}

func crawl(link string, base *url.URL) {
    if visited[link] { // check if the link has already been visited
        return 
    }

    visited[link] = true

    resp, err := http.Get(link)
    if err != nil {
        fmt.Println("Error fetching:", link)
        return
    }
    defer resp.Body.Close()

    z := html.NewTokenizer(resp.Body)
    for {
        tt := z.Next()
        switch {
        case tt == html.ErrorToken:
            return
        case tt == html.StartTagToken:
            t := z.Token()
            if t.Data == "a" {
                for _, a := range t.Attr {
                    if a.Key == "href" {
                        href := a.Val
                        absolute := toAbsURL(href, base)
                        if sameDomain(absolute, base) {
                            crawl(absolute, base)
                        }
                    }
                }
            }
        }
    }
}

func toAbsURL(href string, base *url.URL) string {
    u, err := url.Parse(href)
    if err != nil {
        return ""
    }
    return base.ResolveReference(u).String()
}

func sameDomain(u string, base *url.URL) bool {
    parsed, err := url.Parse(u)
    if err != nil {
        return false
    }
    return parsed.Host == base.Host
}
