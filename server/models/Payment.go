package models

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       string `json:"amount"`
	Dt           string `json:"dt"`
	Bank         string `json:"bank"`
	DeliveryCost string `json:"delivery_cost"`
	GoodsTotal   string `json:"goods_total"`
	CustomFee    string `json:"custom_fee"`
}
