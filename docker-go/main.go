package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

// Ignore favicon requests
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server hitted")
	mu.Lock()
	counter++
	currentCounter := counter
	mu.Unlock()
	fmt.Fprintf(w, "You are visitor number %d\n", currentCounter)

	subPath := ""

	if len(r.URL.Path) > 1 {
		subPath = string(r.URL.Path[1:2])
	} else {
		subPath = "No subpath"
	}

	fmt.Fprintf(w, "First character of path is: %s\n", subPath)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}
