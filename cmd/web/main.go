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
  err := r.ParseForm()
  if err != nil {
    http.Error(w, "Form bad...", http.StatusBadRequest)
  }
  todo := r.PostForm.Get("todo")
  notes := r.PostForm.Get("notes")
  w.Write([]byte(todo + notes))
  http.Redirect(w, r, "/", 200)
}

func main() {
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux := http.NewServeMux()

  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  mux.HandleFunc("/", home)
  mux.HandleFunc("/todo/add", addtodo)

  http.ListenAndServe(":4000", mux)
}
