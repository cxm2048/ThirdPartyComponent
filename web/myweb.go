package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

//func main() {
//    http.HandleFunc("/", handler) // each request calls handler
//    log.Fatal(http.ListenAndServe("localhost:8000", nil))
//}
//
//// handler echoes the Path component of the request URL r.
//func handler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//}

var mu sync.Mutex
var count int

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}
