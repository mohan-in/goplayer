package main

import (
	_ "fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe("localhost:9091", nil)
}
