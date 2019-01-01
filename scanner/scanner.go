package scanner

import (
	"fmt"
	"log"
	"github.com/tjgq/sane"
	"scantastic/file-access"
)

type ScanInstructions struct {
	Filename string `json: filename`
	Foldername string `json: foldername`
}

func Scan(scanInstructions ScanInstructions) (string, []byte, error) {
	if scanInstructions.Filename == "" {
		return "", []byte{}, fmt.Errorf("bad Request: filename was not set")
	}
	if scanInstructions.Foldername == "" {
		return "", []byte{}, fmt.Errorf("bad Request: foldername was not set")
	}
	devs, err := sane.Devices()
	if err != nil {
		return "", []byte{}, fmt.Errorf("could not get a list of devices: %s", err)
	}
	// TODO: should we always pick the first device?
	c, err := sane.Open(devs[0].Name)
	defer c.Close()
	if err != nil {
		return "", []byte{}, fmt.Errorf("could not open a connection to a scanner: %s", err)
	}
	i, err := c.ReadImage()
	if err != nil {
		return "", []byte{}, fmt.Errorf("could not read image from scanner: %s", err)
	}
	path, bytes, err := file_access.WriteImageFile(i, scanInstructions.Filename, scanInstructions.Foldername)
	if err != nil {
		return "", []byte{}, fmt.Errorf("%s", err)
	}
	return path, bytes, nil
}

func Init() {
	err := sane.Init()
	if err != nil {
		log.Fatal("Could not start scanner session", err)
	}
}

func End() {
	sane.Exit()
}