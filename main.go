package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"msvr"
	"net/http"
)

var root string

func init() {
	flag.StringVar(&root, "root", "D:\\", "directory listed in the player when the application starts")
}

func directoryHandler(rw http.ResponseWriter, r *http.Request) {

	var path string
	if path = r.URL.Query().Get("path"); path == "" {
		path = root
	}

	files, err := msvr.Get(path)
	if err != nil {
		fmt.Println(err)
	}

	enc := json.NewEncoder(rw)
	if err := enc.Encode(files); err != nil {
		fmt.Println(err)
	}
}

func mediaFileHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, r.URL.Query().Get("file"))
}

func staticFilesHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, r.URL.Path[1:])
}

func main() {
	flag.Parse()
	http.HandleFunc("/", directoryHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.HandleFunc("/media/", mediaFileHandler)
	http.ListenAndServe("localhost:9091", nil)
}
