package main

import (
    "html/template"
    "log"
    "net/http"
    "strconv"
)

type Book struct {
    ID     int
    Title  string
    Author string
    Year   int
}

var books = []Book{
    {1, "The Go Programming Language", "Alan Donovan", 2015},
    {2, "Clean Code", "Robert C. Martin", 2008},
    {3, "The Pragmatic Programmer", "Andrew Hunt", 1999},
}

func main() {

    // ⚠️ Conflit volontaire : L'autre branche modifie cette ligne
    log.Println("Serveur Go en démarrage...")

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/book", bookHandler)
    http.HandleFunc("/contact", contactHandler)

    log.Println("Serveur lancé sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/home.html"))
    tmpl.Execute(w, books)
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
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

    var selected *Book
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

func contactHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/contact.html"))
    tmpl.Execute(w, nil)
}
