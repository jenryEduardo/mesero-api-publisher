package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)


// Nueva URL del webhook de Discord para revisiones de PR
const discordReviewWebhookURL = "https://discordapp.com/api/webhooks/1349941754798800916/pAs1vcuQMWTeI6oyf0UJ9uGVI-NGk084J8B_FaFNbITVH61BnTK5wItdMpa3OCmEnKZH"

// Función para enviar mensajes a Discord
func SendReviewToDiscord(message string) {
	payload := map[string]string{"content": message}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(discordReviewWebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error al enviar mensaje a Discord:", err)
	} else {
		defer resp.Body.Close()
		log.Println("Mensaje enviado a Discord. Código:", resp.Status)
	}
}

// Endpoint para manejar revisiones de PR
func GithubReviewWebhookHandler(c *gin.Context) {
	var payload github.PullRequestReviewEvent
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := payload.Review
	pr := payload.PullRequest

	if review != nil && pr != nil {
		message := fmt.Sprintf(
			"📝 **Revisión de PR**: [%s](%s)\n👤 Revisor: %s\n📌 Estado: %s\n🗨 Comentario: %s",
			*pr.Title, *pr.HTMLURL, *review.User.Login, *review.State, *review.Body,
		)
		SendReviewToDiscord(message)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook de revisión procesado"})
}