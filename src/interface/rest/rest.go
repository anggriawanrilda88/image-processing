package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handler_mb_app_v1 "github.com/ubersnap-test/src/interface/rest/v1/mobile_app/handlers"
	mobileapp_route "github.com/ubersnap-test/src/interface/rest/v1/mobile_app/routes"
)

const apiName = "ubersnap-test"

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "check_success",
	})
}

func NewRoutes(r *gin.Engine) {
	MobileAppRouteV1(r)
}

func MobileAppRouteV1(r *gin.Engine) {
	// route v1 group
	api := r.Group("/api/v1")

	// handlers
	imageProcessingHandler := handler_mb_app_v1.NewImageProcessingHandler()

	// image processing route
	mobileapp_route.ImageProcessingRoutes(api.Group("/images"), imageProcessingHandler)
}
