package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<font size=\"8\">Hello CodeMash how are you!!</font>")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
