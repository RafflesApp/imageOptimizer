package adapters

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

type UploaderAdapter struct {
	Uploader *s3manager.Uploader
}

func GetUploaderAdapter(client *s3manager.Uploader) *UploaderAdapter {
	return &UploaderAdapter{Uploader: client}
}

func (adapter *UploaderAdapter) UploadImage(file []byte, fileName string) (string, string, error) {
	result, err := adapter.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(file),
	})

	if err != nil {
		log.Printf("Error uploading file %s to S3 bucket: %v", fileName, err)
		return "", "", err
	}
	return result.Location, *result.ETag, nil
}
