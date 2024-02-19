package mappers

type UploadResponse struct {
	ResourceUrl string `json:"resourceUrl"`
	KeyFile     string `json:"key"`
}

func UploadEntityToResponse(fileName string, url string) UploadResponse {
	return UploadResponse{
		KeyFile:     fileName,
		ResourceUrl: url,
	}
}
