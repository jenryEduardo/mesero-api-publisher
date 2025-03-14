package main

import (
	"log"
	"net/http"

	countsRoutes "publisher/src/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	countsRoutes.SetUp(router)

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	
	port := ":8080"
	router.Run(port)
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))

		
}