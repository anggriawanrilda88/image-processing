package mobileapp_route

import (
	"net/http"

	"github.com/ubersnap-test/src/interface/rest/v1/mobile_app/handlers"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "check_success",
	})
}

func ImageProcessingRoutes(r *gin.RouterGroup, handler handlers.ImageProcessingHandler) {
	r.POST("/convert-png-to-jpg", handler.ConvertPNGtoJPEG)
	r.POST("/resize-image", handler.ResizeSpecificImage)
	r.POST("/compress-image", handler.ImageCompress)
	r.GET("/download/:image", handler.Download)
}
