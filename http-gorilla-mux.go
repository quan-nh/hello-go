package main

// go get -u github.com/gorilla/mux
// or use go module

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// handler
func homeHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func productsHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Products!\n"))
}

func productHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Product: %v\n", vars["key"])
}

// midlleware
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}

func main() {
    router := mux.NewRouter()
	router.Use(loggingMiddleware)
	
    // Routes consist of a path and a handler function.
	router.HandleFunc("/", homeHandler)

	// subrouter
	s := router.PathPrefix("/products").Subrouter()
	// "/products/"
	s.HandleFunc("/", productsHandler)
	// "/products/{key}/"
	s.HandleFunc("/{key}", productHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", router))
}
