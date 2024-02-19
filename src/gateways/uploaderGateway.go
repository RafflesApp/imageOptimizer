package gateways

type UploaderGateway interface {
	UploadImage(file []byte, fileName string) (string, string, error)
}

var uploaderGateway UploaderGateway

func SetUploaderGateway(gateway UploaderGateway) {
	uploaderGateway = gateway
}

func UploadImage(file []byte, fileName string) (string, string, error) {
	return uploaderGateway.UploadImage(file, fileName)
}
