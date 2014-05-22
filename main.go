package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type FileList struct {
	Name  string
	IsDir bool
}

var workspaces []string

func init() {
	workspaces = make([]string, 10)
	workspaces[0] = "D:\\music"
}

func handler(rw http.ResponseWriter, req *http.Request) {
	os.Chdir(workspaces[0])
	files, err := ioutil.ReadDir(workspaces[0])
	if err != nil {

	}
	var fileList = make([]FileList, len(files))
	for i, file := range files {
		fileList[i].Name = file.Name()
		fileList[i].IsDir = file.IsDir()
	}

	f, err := json.Marshal(fileList)
	if err != nil {

	}
	fmt.Fprint(rw, f)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9091", nil)
}
