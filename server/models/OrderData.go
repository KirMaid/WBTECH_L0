package models

type OrderData struct {
	Order    Order    `json:"order"`
	Delivery Delivery `json:"delivery"`
	Payment  Payment  `json:"payment"`
	Items    []Item   `json:"items"`
}
