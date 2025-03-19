package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

const discordWebhookURLS = "https://discordapp.com/api/webhooks/1349941754798800916/pAs1vcuQMWTeI6oyf0UJ9uGVI-NGk084J8B_FaFNbITVH61BnTK5wItdMpa3OCmEnKZH"

// SendMessageToDiscord envía un mensaje al canal de Discord
func SendMessageToDiscord(message string) {
	payload := map[string]string{"content": message}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(discordWebhookURLS, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("❌ Error al enviar mensaje a Discord:", err)
		return
	}
	defer resp.Body.Close()
	log.Println("✅ Mensaje enviado a Discord. Código:", resp.Status)
}

// GithubReviewWebhookHandler maneja los eventos de revisión de PR de GitHub
func GithubReviewWebhookHandler(c *gin.Context) {
	var payload github.PullRequestReviewEvent

	// Leer el cuerpo de la solicitud
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("❌ Error al leer el cuerpo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el cuerpo"})
		return
	}
	log.Println("🔹 Payload recibido:", string(body))

	// Volver a leer el cuerpo para el unmarshalling
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Decodificar el JSON
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println("❌ Error al decodificar JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar que el payload contiene datos válidos
	if payload.Review == nil || payload.PullRequest == nil {
		log.Println("❌ Payload inválido: faltan datos de revisión o PR")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido"})
		return
	}

	// Construir el mensaje
	message := fmt.Sprintf(
		"📝 **Revisión de PR**: [%s](%s)\n👤 **Revisor**: %s\n📌 **Estado**: %s\n🗨 **Comentario**: %s",
		*payload.PullRequest.Title, *payload.PullRequest.HTMLURL, *payload.Review.User.Login, *payload.Review.State, *payload.Review.Body,
	)

	// Enviar mensaje a Discord
	SendMessageToDiscord(message)
	log.Println("✅ Notificación enviada a Discord:", message)

	c.JSON(http.StatusOK, gin.H{"message": "Webhook de revisión procesado"})
}
