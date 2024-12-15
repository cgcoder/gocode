package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Args struct {
	port     uint16
	username string
	pass     string
}

type AppServer struct {
	args Args
}

type LoginServer AppServer

func (a *Args) isAuthRequired() bool {
	return a.username != "" && a.pass != ""
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
func isValidCookie(cookie *http.Cookie, args *Args) bool {
	value := cookie.Value
	cred := fmt.Sprintf("%s:%s", args.username, args.pass)
	base64Cred := base64.StdEncoding.EncodeToString([]byte(cred))

	return value == base64Cred
}

func checkSessionAndRedirect(writer http.ResponseWriter, request *http.Request, args *Args) bool {
	if args.isAuthRequired() {
		sessionCookie, err := request.Cookie("session")
		if err != nil || !isValidCookie(sessionCookie, args) {
			http.Redirect(writer, request, "/login?status=auth", http.StatusFound)
			return false
		} else {
			return true
		}
	}
	return true
}

func okStartHtml(writer http.ResponseWriter, title string) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "<html><head><title>%s</title></head><body>", title)
}

func okEndHtml(writer http.ResponseWriter) {
	fmt.Fprintf(writer, "</body></html>")
}

func (s *AppServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if !checkSessionAndRedirect(writer, request, &s.args) {
		return
	}

	path := request.URL.Query().Get("path")
	if path == "" {
		path = "."
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		writeError(writer, http.StatusInternalServerError)
		return
	}
	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			writeError(writer, http.StatusInternalServerError)
			return
		}

		okStartHtml(writer, path)
		fmt.Fprint(writer, "<ul>")
		for _, entry := range entries {
			writeFile(writer, path, entry)
		}
		fmt.Fprint(writer, "</ul>")
		okEndHtml(writer)
	} else {
		ext := filepath.Ext(path)
		mimeType := getMimeTypeByExtension(ext)
		writer.Header().Set("Content-Type", mimeType)
		http.ServeFile(writer, request, path)
	}
}

func writeFile(writer http.ResponseWriter, path string, entry os.DirEntry) {
	suffix := ""
	if entry.IsDir() {
		suffix = "[d]"
	}
	fmt.Fprintf(writer, "<li><a href='?path=%s/%s'>%s</a>%s</li>", path, entry.Name(), entry.Name(), suffix)
}

func writeError(writer http.ResponseWriter, status int) {
	okStartHtml(writer, "Error")
	fmt.Fprintf(writer, "<h1>Error retrieve files.</h1>")
	defer okEndHtml(writer)
}

func (s *LoginServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		username := request.FormValue("username")
		password := request.FormValue("password")
		if username == s.args.username && password == s.args.pass {
			cred := fmt.Sprintf("%s:%s", s.args.username, s.args.pass)
			base64Cred := base64.StdEncoding.EncodeToString([]byte(cred))
			cookie := http.Cookie{
				Name:  "session",
				Value: base64Cred,
			}
			http.SetCookie(writer, &cookie)
			http.Redirect(writer, request, "/", http.StatusFound)
		} else {
			http.Redirect(writer, request, "/login?status=failed", http.StatusFound)
		}
	} else if request.Method == "GET" {
		okStartHtml(writer, "Login")
		defer okEndHtml(writer)
		status := request.URL.Query().Get("status")
		if status == "failed" {
			fmt.Fprintf(writer, "<span style='color: red;'>Login failed. Please try again.</span>")
		} else if status == "auth" {
			fmt.Fprintf(writer, "<span style='color: red;'>Authentication required.</span>")
		}
		fmt.Fprintf(writer, "<form method='POST'>")
		fmt.Fprintf(writer, "<input type='text' name='username' placeholder='Username' />")
		fmt.Fprintf(writer, "<input type='password' name='password' placeholder='Password' />")
		fmt.Fprintf(writer, "<input type='submit' value='Login' />")
		fmt.Fprintf(writer, "</form>")
	} else {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	args := getArgs()
	fmt.Printf("Starting fileserver on port : %d\n", args.port)
	server := AppServer{args: args}

	mux := http.NewServeMux()
	mux.Handle("/", &server)
	mux.Handle("/login", &LoginServer{args: args})
	srv := &http.Server{
		Handler: mux,
		Addr:    fmt.Sprintf(":%d", args.port),
	}
	srv.ListenAndServe()
	fmt.Printf("Good Bye!")
}

func getMimeTypeByExtension(extension string) string {
	switch extension {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	case ".png":
		return "image/png"
	case ".jpg":
		return "image/jpg"
	case ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".ico":
		return "image/x-icon"
	case "txt":
	case "cmd":
	case "sh":
	case "bat":
		return "text/plain"
	}

	return "application/octet-stream"
}
