# Mini Server Hosting: Hands-On Introduction to Core Hosting Concepts

This is a minimal Go web server project designed to explore and understand basic web hosting concepts using [Render](https://render.com/) as a free shared hosting platform.

See also -> [Blog about this project](https://snkzt.github.io/posts/test-share-hosting)

## Project Overview

This project helps learning how to:

- Build a basic HTTP server in Go
- Host the server on Render's shared infrastructure
- Understand how reverse proxies, HTTPS, and automated deployment work IRL
- See Render manage TLS and serve your app securely over HTTPS

## Project Structure
```
/test-share-hosting
├── .gitignore
├── go.mod
├── main.go # The Go web server
└── README.md # This file
```

## How to Run Locally
```
go run main.go
```
The server will start at `http://localhost:10000`.


## How to Deploy on Render

1. Push this code to a GitHub repo
2. Connect the repo to a new Render Web Service
3. Render will:
   - Build your app
   - Detect the port from main.go
   - Expose a public HTTPS URL automatically

> Note: You don’t need to use `http.ListenAndServeTLS` in this project because Render manages HTTPS externally. In a self-hosted setup, you’d need to use `ListenAndServeTLS` for encryption.

© 2025 – snkzt