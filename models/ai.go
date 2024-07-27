package models

type AIRequest struct {
	Input string `json:"input"`
}

type AIResponse struct {
	Type   int    `json:"type"`
	Output string `json:"output"`
}