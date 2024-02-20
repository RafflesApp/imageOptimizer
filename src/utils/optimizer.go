package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
)

const Quality = 70

func OptimizeImage(file []byte, filename string) ([]byte, string, error) {
	var format = GetFileExtension(filename)
	var newImage bytes.Buffer
	switch format {
	case ".jpeg", ".jpg":
		errEncoding := decodeJPEG(&newImage, file)
		if errEncoding != nil {
			return nil, "", errEncoding
		}
	case ".png":
		errEncoding := decodePNG(&newImage, file)
		if errEncoding != nil {
			return nil, "", errEncoding
		}
	default:
		return nil, "", fmt.Errorf("image format %s not supported", format)
	}

	newFileName := fmt.Sprintf("%s.%s", generateNewFilename(), format)

	return newImage.Bytes(), newFileName, nil
}

func decodePNG(buffer *bytes.Buffer, pngData []byte) error {
	img, err := png.Decode(bytes.NewReader(pngData))
	if err != nil {
		log.Printf("Error decoding PNG file: %s", err)
		return err
	}

	return reduceImageQuality(buffer, img)
}

func decodeJPEG(buffer *bytes.Buffer, jpegData []byte) error {
	img, err := jpeg.Decode(bytes.NewReader(jpegData))
	if err != nil {
		log.Printf("Error decodding JPEG image: %v", err)
		return err
	}

	return reduceImageQuality(buffer, img)
}

func reduceImageQuality(buffer *bytes.Buffer, image image.Image) error {
	errEncoding := jpeg.Encode(buffer, image, &jpeg.Options{Quality: Quality})
	if errEncoding != nil {
		log.Printf("Error encoding JPEG image: %v", errEncoding)
		return errEncoding
	}
	return nil
}
