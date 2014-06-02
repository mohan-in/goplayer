package main

import (
	"encoding/json"
	"fmt"
	"msvr"
	"net/http"
)

var workspaces []string

func init() {
	workspaces = make([]string, 10)
	workspaces[0] = "D:\\"
}

func directoryHandler(rw http.ResponseWriter, r *http.Request) {

	path := ""
	if path = r.URL.Query().Get("path"); path == "" {
		path = workspaces[0]
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
	http.HandleFunc("/", directoryHandler)
	http.HandleFunc("/static/", staticFilesHandler)
	http.HandleFunc("/media/", mediaFileHandler)
	http.ListenAndServe("localhost:9091", nil)
}
