package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Handle incoming HTTP requests
// http.Request is a struct that can be quite large and passing it by value
// would be inefficient in terms of memory and performance hence passed by reference
// this will allow handler function to access and modify the request efficiently
func handler(w http.ResponseWriter, r *http.Request) {
	// r (http.Request) object contains all the information about
	// the incoming HTTP request and we use the information with logging
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	// w (http.ResponseWriter) is an interface that allows us
	// to send a response to the client
	fmt.Fprintln(w, "Hello from Go on Render")
}

func main() {
	// http.HandleFunc ties a URL pattern (e.g., "/") to a specific handler function
	// "/" means this handler will respond to all requests that hit the root URL of the server
	// Whenever a client makes a request to "/", this handler func will be called and process the request
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	// Conventional message logging indicating the server is about to start
	// log is a simple way to log information to the console, adding a timestamp automatically
	log.Println("Server is starting on port 10000")
	// Starts the HTTP server and listens for incoming requests
	// This tells the server to listen on port 10000 for HTTP requests
	// ":10000" means listen on port 10000 on all available network interfaces
	// nil indicates that the handler defined ealier with http.HandleFunc will be used
	log.Fatal(http.ListenAndServe(":10000", nil))
	// If server fails to start, it logs the error and stops the programme
	// This is conventional and the execution should be stopped at this step
	// If no error, nothing happen and process will go on

}
