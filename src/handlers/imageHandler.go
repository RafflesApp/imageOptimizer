package handlers

import (
	"github.com/RafflesApp/imageOptimizer/src/constants"
	"github.com/RafflesApp/imageOptimizer/src/exceptions"
	"github.com/RafflesApp/imageOptimizer/src/gateways"
	"github.com/RafflesApp/imageOptimizer/src/mappers"
	"github.com/RafflesApp/imageOptimizer/src/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func Upload(context *gin.Context) {
	fileData, exceptionFile := getFile(context)
	if exceptionFile.ErrorCode != 0 {
		context.AbortWithStatusJSON(exceptionFile.ErrorCode, exceptionFile)
		return
	}

	optimizedImage, errOptimizing := utils.OptimizeImage(fileData)
	if errOptimizing != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, exceptions.Response{}.
			Create(http.StatusBadRequest, constants.FileProcessingError))
		return
	}
	url, tag, errUploading := gateways.UploadImage(optimizedImage, utils.GenerateNewFilename())
	if errUploading != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, exceptions.Response{}.
			Create(http.StatusBadRequest, constants.FileUploadError))
		return
	}
	context.JSON(http.StatusOK, mappers.UploadEntityToResponse(url, tag))
}

func getFile(context *gin.Context) ([]byte, exceptions.Response) {
	file, _, errGetting := context.Request.FormFile(constants.FormDataKey)
	if errGetting != nil {
		log.Printf("Error getting file from request: %s", errGetting)
		return []byte{}, exceptions.Response{}.
			Create(http.StatusBadRequest, constants.FileNotValid)
	}

	fileBytes, errReading := io.ReadAll(file)
	if errReading != nil {
		log.Printf("Error reading file in buffer: %s", errReading)
		return []byte{}, exceptions.Response{}.
			Create(http.StatusBadRequest, constants.FileProcessingError)
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(file)

	return fileBytes, exceptions.Response{}
}
