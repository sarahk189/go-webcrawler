Web Crawler

Overview

It recursively crawls internal links from a starting domain and returns a list of all unique pages in clean JSON format.

Tech Used
Go (Golang)
- `net/http` – to send HTTP requests and fetch HTML
- `net/url` – to parse and resolve absolute/relative URLs
- `golang.org/x/net/html` – to tokenize and parse HTML for `<a href="...">` links

Project Structure

- `main.go` – Starts the server and registers the `/pages` endpoint
- `crawler.go` – Contains the crawling logic:
  - Tracks visited pages
  - Parses and resolves links
  - Validates links belong to the same domain
- `go.mod`, `go.sum` – Module management
