package thumbify

import (
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func ThisImage(filename string) (thumbBytes []byte, err error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(filename); err != nil {
		return nil, fmt.Errorf("could not read stored image file: %s", err)
	}

	h, w, err := mw.GetSize()
	if err != nil {
		return nil, fmt.Errorf("could not get initial image size: %s", err)
	}

	nh := uint(h / 10)
	nw := uint(w / 10)

	if err = mw.ThumbnailImage(nh, nw); err != nil {
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
