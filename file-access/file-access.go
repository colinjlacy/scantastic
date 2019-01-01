package file_access

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"scantastic/thumbify"
)

// TODO: make env var
const basePath = "~/Documents/scanned/"

func WriteImageFile(i image.Image, filename string, filepath string) (fullFilePath string, thumbBytes []byte, err error) {
	fullpath := basePath + filepath
	e, err := pathExists(fullpath)
	if err != nil {
		return "", []byte{}, fmt.Errorf("could not determine the existence of the desired location: %s", err)
	}
	if !e {
		err = createPath(fullpath)
		if err != nil {
			return "", []byte{}, fmt.Errorf("could not create the folder path to the desired location: %s", err)
		}
	}
	fullFilePath = fullpath + "/" + filename + ".jpg"
	f, err := os.Create(fullFilePath)
	if err != nil {
		return "", []byte{}, fmt.Errorf("could not create a file at the desired location: %s", err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, i, nil); err != nil {
		return "", []byte{}, fmt.Errorf("could not write jpeg data to file: %s", err)
	}
	thumbBytes, err = thumbify.ThisImageFile(f)
	if err != nil {
		fmt.Println("could not generate thumbnail byte slice for the scanned image %s: %s", filename, err)
		err = nil
	}
	thumbFilePath := fullpath + "/thumbs/" + filename + ".jpg"
	t, err := os.Create(thumbFilePath)
	if err != nil {
		fmt.Println("could not create a thumbnail for the scanned image %s: %s", filename, err)
		err = nil
	} else {
		defer t.Close()
		_, err = t.Write(thumbBytes)
		if err != nil {
			fmt.Println("could not write thumbnail to file for image %s: %s", filename, err)
			err = nil
		}
	}

	return fullFilePath, thumbBytes,nil
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