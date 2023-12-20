package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Params: ", r.URL, r.Method)
		h.ServeHTTP(w, r)
	})
}

func PongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("pong!"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := mux.NewRouter()
	// use the middleware
	mux.Use(Middleware)

	// set the pong handler
	mux.HandleFunc("/ping", PongHandler)

	// get the port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
		fmt.Println("Defaulting to port " + port)
	}
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		fmt.Errorf("Server failed to start %s", err)
	}
}
