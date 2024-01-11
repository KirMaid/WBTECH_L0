package models

type Payment struct {
	OrderUid     string  `json:"order_uid"`
	Transaction  string  `json:"transaction"`
	RequestId    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       float64 `json:"amount"`
	Dt           int64   `json:"dt"`
	Bank         string  `json:"bank"`
	DeliveryCost float64 `json:"delivery_cost"`
	GoodsTotal   float64 `json:"goods_total"`
	CustomFee    float64 `json:"custom_fee"`
}
