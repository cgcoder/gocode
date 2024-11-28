package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cgcoder/gocode/proxy/config"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	io.WriteString(w, "Hello, World!")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!")
	io.WriteString(w, "Home Page!")
	config.SayHello()
}

func main() {
	const PORT = ":6565"
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	fmt.Println("Server listing ", PORT)
	http.ListenAndServe(PORT, nil)
}
