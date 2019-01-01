package thumbify

import (
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
	"os"
)

func ThisImageFile(f *os.File) (thumbBytes []byte, err error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImageFile(f); err != nil {
		return nil, fmt.Errorf("could not read stored image file: %s", err)
	}

	if err != nil {
		return nil, fmt.Errorf("could not get initial image size: %s", err)
	}

	h := uint(200)
	w := uint(200)

	if err = mw.ThumbnailImage(h, w); err != nil {
		return nil, fmt.Errorf("could not resize image to thumbnail: %s", err)
	}

	thumbBytes = mw.GetImagesBlob()
	return thumbBytes, err
}

func Start() {
	imagick.Initialize()
}

func End() {
	imagick.Terminate()
}
