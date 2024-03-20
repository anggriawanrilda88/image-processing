package transformers

import (
	"github.com/ubersnap-test/src/domain/entities"
)

type imageProcessingDownload struct {
	LinkDownload string `json:"linkDownload"`
}

// ResponseListHandler, response format for list data
func ImageProcessingDownload(data *entities.ImageProcessing) *imageProcessingDownload {
	return &imageProcessingDownload{
		LinkDownload: data.GetOutputPath(),
	}
}
