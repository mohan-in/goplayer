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

func directoryHandler(rw http.ResponseWriter, r *http.Request) {

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
