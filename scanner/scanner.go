package scanner

import (
	"fmt"
	"github.com/tjgq/sane"
	"log"
	"os"
	"image/jpeg"
)

type ScanInstructions struct {
	Filename string `json: filename`
	Foldername string `json: foldername`
}

func Scan(scanInstructions ScanInstructions) error {
	if scanInstructions.Filename == "" {
		return fmt.Errorf("bad Request: filename was not set")
	}
	if scanInstructions.Foldername == "" {
		return fmt.Errorf("bad Request: foldername was not set")
	}
	if err := sane.Init(); err != nil {
		return fmt.Errorf("could not start scanner session: %s", err)
	}
	c, err := sane.Open("")
	defer c.Close()
	if err != nil {
		return fmt.Errorf("could not open a connection to a scanner: %s", err)
	}
	i, err := c.ReadImage()
	if err != nil {
		return fmt.Errorf("could not read image from scanner: %s", err)
	}
	f, err := os.Create("./" + scanInstructions.Filename + ".jpg")
	if err != nil {
		return fmt.Errorf("could not create a file at the desired location: %s", err)
	}
	defer f.Close()
	jpeg.Encode(f, i, nil)
	return nil
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