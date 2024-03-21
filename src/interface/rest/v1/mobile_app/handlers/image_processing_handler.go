package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/image-processing/src/app/usecases"
	"github.com/image-processing/src/interface/rest/response"
	"github.com/image-processing/src/interface/rest/v1/mobile_app/requests"
	"github.com/image-processing/src/interface/rest/v1/mobile_app/transformers"
)

type ImageProcessingHandler interface {
	ConvertPNGtoJPEG(c *gin.Context)
	ResizeSpecificImage(c *gin.Context)
	ImageCompress(c *gin.Context)
	Download(c *gin.Context)
}

type imageProcessingHandler struct {
	usecase usecases.ImageProcessingAppInterface
}

func NewImageProcessingHandler() ImageProcessingHandler {
	usecase := usecases.NewImageProcessingUsecase()
	return &imageProcessingHandler{
		usecase: usecase,
	}
}

func (h *imageProcessingHandler) ConvertPNGtoJPEG(c *gin.Context) {
	req := requests.ConvertPNGtoJPEG{}
	dto, err := req.Validate(c)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	imageProcess, err := h.usecase.ConvertPNGtoJPEG(c, dto)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	response.ResponseJSON(c, "Successfully convert PNG to JPEG", transformers.ImageProcessingDownload(imageProcess))
}

func (h *imageProcessingHandler) ResizeSpecificImage(c *gin.Context) {
	req := requests.ResizeSpecificImage{}
	dto, err := req.Validate(c)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	imageProcess, err := h.usecase.ResizeSpecificImage(c, dto)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	response.ResponseJSON(c, "Successfully convert PNG to JPEG", transformers.ImageProcessingDownload(imageProcess))
}

func (h *imageProcessingHandler) ImageCompress(c *gin.Context) {
	req := requests.ImageCompress{}
	dto, err := req.Validate(c)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	imageProcess, err := h.usecase.ImageCompress(c, dto)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	response.ResponseJSON(c, "Successfully convert PNG to JPEG", transformers.ImageProcessingDownload(imageProcess))
}

func (h *imageProcessingHandler) Download(c *gin.Context) {
	req := requests.ImageDownload{}
	dto, err := req.Validate(c)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	imageProcess, err := h.usecase.Download(c, dto)
	if err != nil {
		response.ErrorHandler(c, err)
		return
	}

	response.ResponseImage(c, dto.Image, "image/jpeg", imageProcess.GetFileContent())
}
