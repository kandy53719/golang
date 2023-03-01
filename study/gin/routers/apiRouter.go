package route

import (
	"golang/study/gin/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRouteInit(route *gin.Engine) {
	var ac = controllers.ApiCintroller{}
	route.GET("/api", ac.Get)
	route.POST("/api", ac.Post)
	route.PUT("/api", ac.Put)
	route.DELETE("/api", ac.Delete)
}
