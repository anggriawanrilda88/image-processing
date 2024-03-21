package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/image-processing/src/app/dto"
	"github.com/image-processing/src/domain/entities"
)

type imageProcessingApp struct {
}

type ImageProcessingAppInterface interface {
	ConvertPNGtoJPEG(ctx *gin.Context, dto *dto.ConvertPNGtoJPEGDTO) (*entities.ImageProcessing, error)
	ResizeSpecificImage(ctx *gin.Context, dto *dto.ResizeSpecificImage) (*entities.ImageProcessing, error)
	ImageCompress(ctx *gin.Context, dto *dto.ImageCompress) (*entities.ImageProcessing, error)
	Download(ctx *gin.Context, dto *dto.ImageDownload) (*entities.ImageProcessing, error)
}

func NewImageProcessingUsecase() ImageProcessingAppInterface {
	return &imageProcessingApp{}
}

func (u *imageProcessingApp) ConvertPNGtoJPEG(ctx *gin.Context, dto *dto.ConvertPNGtoJPEGDTO) (*entities.ImageProcessing, error) {
	return entities.CreateConverPNGtoJPEG(
		dto.Image,
		dto.Header,
		ctx.Request.Host,
	)
}

func (u *imageProcessingApp) ResizeSpecificImage(ctx *gin.Context, dto *dto.ResizeSpecificImage) (*entities.ImageProcessing, error) {
	return entities.ResizeSpecificImage(
		dto.Image,
		dto.Header,
		dto.Width,
		dto.Height,
		ctx.Request.Host,
	)
}

func (u *imageProcessingApp) ImageCompress(ctx *gin.Context, dto *dto.ImageCompress) (*entities.ImageProcessing, error) {
	return entities.ImageCompress(
		dto.Image,
		dto.Header,
		dto.ImageQuality,
		ctx.Request.Host,
	)
}

func (u *imageProcessingApp) Download(ctx *gin.Context, dto *dto.ImageDownload) (*entities.ImageProcessing, error) {
	return entities.Download(
		dto.Image,
	)
}
