package routes

import (
	"api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()
	urlRoot := "/tweet"

	v1 := router.Group("/v1")
	{
		v1.GET(urlRoot, tweetController.FindAll)
		v1.POST(urlRoot, tweetController.Create)
		v1.DELETE(urlRoot+"/:id", tweetController.Delete)
	}

	return v1
}
