package main

import (
	id3 "github.com/gocode/go-id3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type MediaFile struct {
	Name    string
	IsDir   bool
	AbsPath string
	Size    float64
	IsAudio bool
	IsVideo bool
	ID3Name string
	Artist  string
	Album   string
	Length  string
}

func getFiles(path string) ([]MediaFile, error) {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files = make([]MediaFile, 0, len(fileInfo))
	var f MediaFile

	for _, fi := range fileInfo {
		switch strings.ToLower(filepath.Ext(fi.Name())) {
		case ".mp3", ".ogg":
			var fd, err = os.Open(filepath.Join(path, fi.Name()))
			if err != nil {
				return nil, err
			}
			id, err := id3.Read(fd)
			if err != nil {
				id = &id3.File{}
			}
			f = MediaFile{
				Name:    fi.Name(),
				IsDir:   fi.IsDir(),
				AbsPath: filepath.Join(path, fi.Name()),
				Size:    float64(fi.Size()) / (1024 * 1024),
				IsAudio: true,
				ID3Name: id.Name,
				Artist:  id.Artist,
				Album:   id.Album,
				Length:  id.Length,
			}
			files = append(files, f)
		case "", ".mp4":
			f = MediaFile{
				Name:    fi.Name(),
				IsDir:   fi.IsDir(),
				AbsPath: filepath.Join(path, fi.Name()),
				Size:    float64(fi.Size()) / (1024 * 1024),
			}
			files = append(files, f)
		}
	}
	return files[:len(files)], nil
}
