package controllers
import (
	"log"
	"net/http"
	"publisher/src/application"
	"publisher/src/infraestructure"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) 
	if err != nil {
		log.Println("Error al convertir el ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido", "details": err.Error()})
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NuevoObtenerPedido(repo)

	pedido, err := useCase.Execute(id)
	if err != nil {
		log.Println("Error al obtener el pedido:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pedido": pedido})
}
