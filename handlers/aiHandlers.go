package handlers

import (
	"bytes"
	"encoding/json"
	"geniva/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CallAIHandler(c *gin.Context) {
	var request struct {
		InputString string `json:"input_string"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	aiRequest := models.AIRequest{Input: request.InputString}
	jsonValue, err := json.Marshal(aiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request"})
		return
	}

	aiEndpoint := os.Getenv("AI_ENDPOINT")
	if aiEndpoint == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI endpoint not set"})
		return
	}

	resp, err := http.Post(aiEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call AI service"})
		return
	}
	defer resp.Body.Close()

	var aiResponse models.AIResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode AI response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"type": aiResponse.Type, "output": aiResponse.Output})
}
