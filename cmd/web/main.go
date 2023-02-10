package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  files := []string{
    "./ui/html/base.tmpl.html",
    "./ui/html/pages/home.tmpl.html",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
    return
  }
  err = ts.ExecuteTemplate(w, "base", nil)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
  }

}

func addtodo(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Add todo form..."))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/todo/add", addtodo)
  http.ListenAndServe(":4000", mux)
}
