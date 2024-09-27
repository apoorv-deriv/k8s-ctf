package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	isHealthy = false
	bigSlice  []byte
)

func main() {
	// Allocate 200MB after 3 minutes
	go func() {
		time.Sleep(3 * time.Minute)
		bigSlice = make([]byte, 200*1024*1024) // 200MB
		fmt.Println("Allocated 200MB")
	}()

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/healthz", handleHealth)
	http.HandleFunc("/ready", handleReady)

	fmt.Println("Server starting on :8000")
	http.ListenAndServe(":8000", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if !isHealthy {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Unhealthy")
		return
	}
	fmt.Fprintf(w, "Healthy")
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	if !isHealthy {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Not Ready")
		return
	}
	fmt.Fprintf(w, "Ready")
}
