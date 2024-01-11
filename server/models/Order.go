package models

type Order struct {
	OrderUid          string `json:"order_uid"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	Shardkey          string `json:"shardkey"`
	SmId              string `json:"sm_id"`
	DateCreated       string `json:"date_created"`
	OofShard          string `json:"oof_shard"`
}
