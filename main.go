package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type FileList struct {
	Name    string
	IsDir   bool
	AbsPath string
}

var workspaces []string

func init() {
	workspaces = make([]string, 10)
	workspaces[0] = "D:\\music"
}

func handler(rw http.ResponseWriter, r *http.Request) {

	path := ""
	if path = r.URL.Query().Get("path"); path == "" {
		path = workspaces[0]
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	var fileList = make([]FileList, len(files))
	for i, file := range files {
		fileList[i].Name = file.Name()
		fileList[i].IsDir = file.IsDir()
		fileList[i].AbsPath = filepath.Join(path, file.Name())
	}

	enc := json.NewEncoder(rw)
	if err := enc.Encode(fileList); err != nil {
		fmt.Println(err)
	}
}

func fileServer(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, r.URL.Query().Get("file"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/mediafile/", fileServer)
	http.ListenAndServe("localhost:9091", nil)
}
