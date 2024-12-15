package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Args struct {
	port     uint16
	username string
	pass     string
}

func getArgs() Args {
	var args Args = Args{
		port: 8899,
	}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-p" {
			if i+1 < len(os.Args) {
				i, err := strconv.Atoi(os.Args[i+1])
				if err != nil {
					panic("Invalid port number")
				} else {
					args.port = uint16(i)
				}
			}
		} else if os.Args[i] == "-u" {
			if i+2 < len(os.Args) {
				args.username = os.Args[i+1]
				args.pass = os.Args[i+2]
			} else {
				panic("-u requires username and password")
			}
		}
	}
	return args
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World!")
}

func main() {
	args := getArgs()
	fmt.Printf("Starting fileserver on port : %d\n", args.port)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))
	srv := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%d", args.port),
	}
	srv.ListenAndServe()
	fmt.Printf("Good Bye!")
}
