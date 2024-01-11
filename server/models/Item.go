package models

type Item struct {
	ChrtID      string `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       string `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        string `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  string `json:"total_price"`
	NMID        string `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      string `json:"status"`
}
