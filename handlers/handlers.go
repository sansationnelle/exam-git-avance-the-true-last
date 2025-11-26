package handlers

import (
	"html/template"
	"net/http"
	"strconv"
    "main/data"
)

var books = []data.Book{
    {1, "The Go Programming Language", "Alan Donovan", 2015},
    {2, "Clean Code", "Robert C. Martin", 2008},
    {3, "The Pragmatic Programmer", "Andrew Hunt", 1999},
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")

    if idStr == "" {
        http.Error(w, "ID requis", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID invalide", http.StatusBadRequest)
        return
    }

    var selected *data.Book
    for _, b := range books {
        if b.ID == id {
            selected = &b
        }
    }

    if selected == nil {
        http.Error(w, "Livre introuvable", http.StatusNotFound)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/book.html"))
    tmpl.Execute(w, selected)
}
