package main

import (
    "net/http"
)

type UserData struct {
    Username string
    UserId   string
    Email    string
}

type Context struct {
    Response http.ResponseWriter
    Request  *http.Request
    User     *UserData 
}

func context(writer http.ResponseWriter, request *http.Request) *Context {
    return &Context {
        Response: writer,
        Request:  request,
    }
}
