package file_access

import (
	"archivalist/manifest-reader"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"
	"scantastic/thumbify"
	"time"
)

// TODO: make env var
var home, _ = homedir.Dir()
var basePath = home + "/Documents/scanned/"

func WriteImageFile(i image.Image, filename string, filepath string) (fullFilePath string, thumbBase64 string, err error) {
	fullpath := basePath + filepath
	e, err := pathExists(fullpath)
	if err != nil {
		return "", "", fmt.Errorf("could not determine the existence of the desired location: %s", err)
	}
	if !e {
		err = createPath(fullpath)
		if err != nil {
			return "", "", fmt.Errorf("could not create the folder path to the desired location: %s", err)
		}
	}
	fullFilePath = fullpath + "/" + filename + ".jpg"
	f, err := os.Create(fullFilePath)
	if err != nil {
		return "", "", fmt.Errorf("could not create a file at the desired location: %s", err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, i, nil); err != nil {
		return "", "", fmt.Errorf("could not write jpeg data to file: %s", err)
	}
	thumbBase64, err = thumbify.ThisImageFile(fullpath, filename)
	if err != nil {
		fmt.Println(err)
		err = nil
	}

	return fullFilePath, thumbBase64, nil
}

func WriteSummaryFile(foldername string, prettyName string) error {
	s := path.Join(basePath, foldername)
	info, err := os.Stat(s)
	if err != nil {
		return fmt.Errorf("could not find the specified folder")
	}
	if !info.IsDir() {
		return fmt.Errorf("the specified folder is not actually a folder")
	}
	files, err := ioutil.ReadDir(s)
	if !info.IsDir() {
		return fmt.Errorf("could not read the specified folder")
	}
	currentTime := time.Now()
	expires := currentTime.AddDate(0,0,7)
	summary := manifest_reader.JobSummary{foldername, prettyName, len(files), currentTime, expires}
	jsonData, _ := json.Marshal(summary)
	filepath := path.Join(s, "manifest.json")
	err = ioutil.WriteFile(filepath, jsonData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not write manifest file")
	}
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func createPath(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
