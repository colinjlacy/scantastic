package file_access

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// TODO: make env var
var basePath = "~/Documents/scanned/"

func WriteImageFile(i image.Image, filename string, filepath string) (string, error) {
	fullpath := basePath + filepath
	e, err := pathExists(fullpath)
	if err != nil {
		return "", fmt.Errorf("could not determine the existence of the desired location: %s", err)
	}
	if !e {
		err = createPath(fullpath)
		if err != nil {
			return "", fmt.Errorf("could not create the folder path to the desired location: %s", err)
		}
	}
	fullFilePath := fullpath + "/" + filename + ".jpg"
	f, err := os.Create(fullFilePath)
	if err != nil {
		return "", fmt.Errorf("could not create a file at the desired location: %s", err)
	}
	defer f.Close()
	jpeg.Encode(f, i, nil)
	return fullFilePath, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func createPath(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}