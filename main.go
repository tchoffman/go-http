package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
	io.WriteString(w, fmt.Sprintf("Oh Hi! You've requested %s\n", r.URL.Path))
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
	io.WriteString(w, fmt.Sprintf("Ya Hallo? You've requested %s\n", r.URL.Path))
}

func getCountUp(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /countup request\n")
	io.WriteString(w, "Counting Up!\n")
	countUp(w, 50)

}

func countUp(w http.ResponseWriter, n int) {
	for i := 0; i <= n; i++ {
		io.WriteString(w, fmt.Sprintf("Count: %d\n", i))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/countup", getCountUp)

	fmt.Println("Server is starting on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed\n")
	} else if err != nil {
		fmt.Printf("Error Starting Server: %s\n", err)
		os.Exit(1)
	}
}
