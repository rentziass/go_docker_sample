package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Altrimenti gabri non ci crede")
	fmt.Println("Starting off on port 3000")
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from a Golang app inside a container! It's nice and warm in here! 🍻")
}
