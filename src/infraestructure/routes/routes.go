package routes

import (
	"publisher/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine){
		routes:=router.Group("/pedido")


		routes.POST("/",controllers.Guardar)

}