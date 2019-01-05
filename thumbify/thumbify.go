package thumbify

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
	"os"
)

func ThisImageFile(folderpath , filename string) (thumbString string, err error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(folderpath + "/" + filename + ".jpg"); err != nil {
		return "", fmt.Errorf("could not read stored image file: %s", err)
	}

	if err != nil {
		return "", fmt.Errorf("could not get initial image size: %s", err)
	}

	h := uint(200)
	w := uint(200)
	thumbFolder := folderpath + "/thumbs"
	if err = os.MkdirAll(thumbFolder, os.ModePerm); err != nil {
		return "", fmt.Errorf("could create thumbnail folder structure: %s", err)
	}

	if err = mw.ThumbnailImage(h, w); err != nil {
		return "", fmt.Errorf("could not resize image to thumbnail: %s", err)
	}

	thumbFilePath := thumbFolder + "/" + filename + ".jpg"
	if err = mw.WriteImage(thumbFilePath); err != nil {
		return "", fmt.Errorf("could not write thumbnail to file: %s", err)
	}

	thumbBytes := mw.GetImagesBlob()
	thumbString = base64.StdEncoding.EncodeToString(thumbBytes)
	return thumbString, err
}

func Start() {
	imagick.Initialize()
}

func End() {
	imagick.Terminate()
}
