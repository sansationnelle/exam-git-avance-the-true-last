package handlers

import (
	"net/http"
    "main/data"
)

var books = []data.Book{
    {1, "The Go Programming Language", "Alan Donovan", 2015},
    {2, "Clean Code", "Robert C. Martin", 2008},
    {3, "The Pragmatic Programmer", "Andrew Hunt", 1999},
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    return
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
    return
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
    return
}