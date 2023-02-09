package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello"))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  http.ListenAndServe(":4000", mux)
}
