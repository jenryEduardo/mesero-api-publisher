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

// URL del webhook de Discord (coloca la tuya aquí)
const discordWebhookURL = "https://discordapp.com/api/webhooks/1349929254573572096/jKxQ52GZeZRR3F9u2QNxcBmF8qsKMDT9a-BVA1ini9mtCo18ib4DSAkEUQ93ff4I7UTq"

// Formato del mensaje para Discord
func SendToDiscord(message string) {
	payload := map[string]string{"content": message}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error al enviar mensaje a Discord:", err)
	} else {
		defer resp.Body.Close()
		log.Println("Mensaje enviado a Discord. Código:", resp.Status)
	}
}

// Maneja los webhooks de GitHub
func GithubWebhookHandler(c *gin.Context) {
	var payload github.PullRequestEvent
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generar mensaje
	pr := payload.PullRequest
	if pr != nil {
		message := fmt.Sprintf("🔔 **Nuevo PR**: [%s](%s)\n👤 Autor: %s\n📌 Estado: %s",
			*pr.Title, *pr.HTMLURL, *pr.User.Login, *payload.Action)
		SendToDiscord(message)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook procesado"})
}