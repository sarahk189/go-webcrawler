package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    "net/url"
)

var visited = make(map[string]bool)

func Crawl(start string) []string {
    visited = make(map[string]bool) // reset
    base, _ := url.Parse(start)
    crawl(start, base)
    
    pages := []string{}
    for page := range visited {
        pages = append(pages, page)
    }
    return pages
}

func crawl(link string, base *url.URL) {
    if visited[link] {
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
