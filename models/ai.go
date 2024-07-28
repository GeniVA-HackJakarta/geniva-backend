package models

type AIRequest struct {
	Query string  `json:"query,omitempty"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

type AIResponse struct {
	Message struct {
		Input       string `json:"input"`
		Output      []int  `json:"output"`
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"message"`
}
