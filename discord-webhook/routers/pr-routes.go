package routers

import (
	"github.com/gin-gonic/gin"
	"webhook/controllers"
)


func SetupRoutesCount(router *gin.Engine) {

	

	{
		router.POST("/webhoo-pr", controllers.GithubWebhookHandler)
		router.POST("/webhook-errors", controllers.GithubReviewWebhookHandler)
	}	
}	


//https://discordapp.com/api/webhooks/1349941754798800916/pAs1vcuQMWTeI6oyf0UJ9uGVI-NGk084J8B_FaFNbITVH61BnTK5wItdMpa3OCmEnKZH