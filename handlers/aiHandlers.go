package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"geniva/models"
	"geniva/repository"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CallAIHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		InputString string `json:"input_string"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := repository.UserRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	modifiedInputString := fmt.Sprintf("Nama: %s, %s", user.Name, request.InputString)
	aiRequest := models.AIRequest{Query: modifiedInputString}
	jsonValue, err := json.Marshal(aiRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request"})
		return
	}
	fmt.Println("AI Request: ", string(jsonValue))

	aiEndpoint := "http://170.64.228.233:8000/generate-recommendation"
	resp, err := http.Post(aiEndpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call AI service"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read AI response"})
		return
	}
	fmt.Println("AI Response Status:", resp.Status)
	fmt.Println("AI Response Body:", string(body))

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": string(body)})
		return
	}

	var aiResponse models.AIResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode AI response"})
		return
	}

	menuIDs := aiResponse.Message.Output
	menus, err := repository.MenuRepo.GetMenusByIDs(menuIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve menus"})
		return
	}

	fmt.Println("AI Response Struct:", aiResponse)

	c.JSON(http.StatusOK, gin.H{
		"description": aiResponse.Message.Description,
		"type":        aiResponse.Message.Type,
		"menus":       menus,
	})
}
