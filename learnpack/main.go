package main

import (
	"fmt"
	"net/http"

	"github.com/gocode/learnpack/server"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s New Text!", request.URL.Path[1:])
}

func preferenceHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Preference Page")
}

func preference(context *Context) {

}

func main() {
	server.InitServer()

	defer server.UnInitServer()
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/preference", preferenceHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
