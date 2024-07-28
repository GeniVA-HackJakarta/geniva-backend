package models

type AIRequest struct {
	Query string  `json:"query,omitempty"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

type TransitStep struct {
	Type     string `json:"type"`
	Distance string `json:"distance"`
	Duration string `json:"duration"`
	Price    string `json:"price"`
}

type Transit struct {
	TotalDistance string        `json:"total_distance"`
	TotalDuration string        `json:"total_duration"`
	Steps         []TransitStep `json:"steps"`
}

type GrabRide struct {
	TotalDistance string `json:"total_distance"`
	TotalDuration string `json:"total_duration"`
	Price         string `json:"price"`
	Type          string `json:"type"`
}

type AIResponse struct {
	Message struct {
		Input       string     `json:"input,omitempty"`
		Output      []int      `json:"output,omitempty"`
		Description string     `json:"description,omitempty"`
		Type        string     `json:"type,omitempty"`
		Transit     *Transit   `json:"transit,omitempty"`
		GrabRide    []GrabRide `json:"grab-ride,omitempty"`
	} `json:"message"`
}
