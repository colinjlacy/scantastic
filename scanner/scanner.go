package scanner

import (
	"fmt"
	//"github.com/tjgq/sane"
	"log"
	"os"
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
	cmd, err := exec.Command("scanimage", scanInstructions.Filename)
	if err != nil {
		return fmt.Errorf("could not intialize command: %s", err)
	}
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("error encountered when executing scan: %s", err)
	}
	//if err := sane.Init(); err != nil {
	//	return fmt.Errorf("could not start scanner session: %s", err)
	//}
	//c, err := sane.Open("")
	//defer c.Close()
	//if err != nil {
	//	return fmt.Errorf("could not open a connection to a scanner: %s", err)
	//}
	//i, err := c.ReadImage()
	//if err != nil {
	//	return fmt.Errorf("could not read image from scanner: %s", err)
	//}
	//i := []byte("hello\ngo\n")
	//f, err := os.Create(scanInstructions.Foldername + "/" + scanInstructions.Filename + ".txt")
	//if err != nil {
	//	return fmt.Errorf("could not create a file at the desired location: %s", err)
	//}
	//defer f.Close()
	//f.Write(i)
	//jpeg.Encode(f, i, nil)
	return nil
}

