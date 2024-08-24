package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Leroy")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
