package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>“Taking your time is important. You should never be in a rush. But, when the world rushes you, don't be rushed. -MarkLee”</h1>")
	
	fmt.Println("Someone hit me!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :80")
	http.ListenAndServe(":80", nil)
}