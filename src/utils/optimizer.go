package utils

import (
	"bytes"
	"fmt"
	"github.com/chai2010/webp"
	"image/jpeg"
)

const NewQuality = 0.7

func OptimizeImage(file []byte) ([]byte, error) {
	img, errDecoding := jpeg.Decode(bytes.NewReader(file))
	if errDecoding != nil {
		fmt.Printf("Error decoding image: %v", errDecoding)
		return []byte{}, errDecoding
	}
	opts := &webp.Options{
		Lossless: false,
		Quality:  NewQuality,
	}

	buf := new(bytes.Buffer)

	if errEncoding := webp.Encode(buf, img, opts); errEncoding != nil {
		fmt.Printf("Error encoding image to webp: %v", errEncoding)
		return []byte{}, errEncoding
	}

	return buf.Bytes(), nil
}
