package main

import (
	"io/ioutil"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

func watchDir(dirName string) ([]string, error) {
	allfiles, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Error("Error in reading directory", err)
		return nil, err
	}

	// Get only files which are written fully
	var files []string
	for _, file := range allfiles {
		// Check if it is a file and have not been modified atleast since last minute
		if file.Mode().IsRegular() && file.ModTime().Add(60*time.Second).Unix() < time.Now().Unix() {
			files = append(files, file.Name())
		}
	}

	var newFiles []string
	for _, file := range files {
		if strings.Contains(file, srcExtension) {
			if !isFileIndexed(file) {
				newFiles = append(newFiles, file)
			}
		}
	}

	return newFiles, nil
}
