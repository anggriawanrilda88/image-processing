package requests

import (
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/ubersnap-test/src/app/dto"
)

// IOrderRequest ...
type ImageProcessingRequest interface {
	Validate(c *gin.Context) (*dto.ConvertPNGtoJPEGDTO, error)
}

type ConvertPNGtoJPEG struct {
	Image      multipart.File
	OutputPath string
}

func (req *ConvertPNGtoJPEG) Validate(c *gin.Context) (*dto.ConvertPNGtoJPEGDTO, error) {
	// Menerima file gambar dari form-data
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		return nil, errors.New("bad request")
	}
	defer file.Close()

	return &dto.ConvertPNGtoJPEGDTO{
		Image:  file,
		Header: header,
	}, nil
}

type ResizeSpecificImage struct {
	Image  multipart.File
	Width  string
	Height string
}

func (req *ResizeSpecificImage) Validate(c *gin.Context) (*dto.ResizeSpecificImage, error) {
	// Menerima file gambar dari form-data
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		return nil, errors.New("bad request")
	}
	defer file.Close()

	req.Image = file
	req.Width = c.PostForm("width")
	req.Height = c.PostForm("height")

	return &dto.ResizeSpecificImage{
		Image:  req.Image,
		Header: header,
		Width:  req.Width,
		Height: req.Height,
	}, nil
}

type ImageCompress struct {
	Image        multipart.File
	ImageQuality string
}

func (req *ImageCompress) Validate(c *gin.Context) (*dto.ImageCompress, error) {
	// Menerima file gambar dari form-data
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		return nil, errors.New("bad request")
	}
	defer file.Close()

	req.Image = file
	req.ImageQuality = c.PostForm("imageQuality")

	return &dto.ImageCompress{
		Image:        req.Image,
		Header:       header,
		ImageQuality: req.ImageQuality,
	}, nil
}

type ImageDownload struct {
	Image string
}

func (req *ImageDownload) Validate(c *gin.Context) (*dto.ImageDownload, error) {
	req.Image = c.Param("image")

	return &dto.ImageDownload{
		Image: req.Image,
	}, nil
}
