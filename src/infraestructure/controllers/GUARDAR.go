package controllers

import (
	"fmt"
	"log"
	"net/http"
	"publisher/src/application"
	"publisher/src/domain"
	"publisher/src/infraestructure"

	"github.com/gin-gonic/gin"
)

func Guardar(c *gin.Context) {

	var pedidp domain.Pedido

	if err := c.ShouldBindJSON(&pedidp); err != nil {
		log.Println(" Error al procesar el JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON", "details": err.Error()})
		return
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NuevoPedido(repo)

	success:=useCase.Execute(pedidp)

	if success==nil{
		fmt.Println("no se pudo ejecutar el caso de uso")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pedido creado con éxito"})


}