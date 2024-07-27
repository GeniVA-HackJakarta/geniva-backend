package models

type Restaurant struct {
	IDZomato           string    `json:"id_zomato" gorm:"primaryKey"`
	Name               string `json:"name"`
	Address            string `json:"address"`
	Zipcode            string `json:"zipcode"`
	Locality           string `json:"locality"`
	City               string `json:"city"`
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
	Cuisines           string `json:"cuisines"`
	Currency           string `json:"currency"`
	AverageCostForTwo  string    `json:"average_cost_for_two"`
	Rating             string `json:"rating"`
	Votes              string    `json:"votes"`
}
