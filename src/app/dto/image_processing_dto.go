package dto

import "mime/multipart"

type ConvertPNGtoJPEGDTO struct {
	Image  multipart.File
	Header *multipart.FileHeader
}

type ResizeSpecificImage struct {
	Image  multipart.File
	Header *multipart.FileHeader
	Width  string
	Height string
}

type ImageCompress struct {
	Image        multipart.File
	Header       *multipart.FileHeader
	ImageQuality string
}

type ImageDownload struct {
	Image string
}
