package main

import (
	database "L0/db"
	"L0/models"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
)

var clientID = "server"
var sc stan.Conn

func setupSubscribe() {
	_, err := sc.QueueSubscribe("order", "queue-group", func(msg *stan.Msg) {
		var orderData models.Order
		err := json.Unmarshal(msg.Data, &orderData)
		if err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}
		err = setOrder(msg.Data)

		if err != nil {
			log.Fatalf("Failed to insert to DB: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to Subscribe: %v", err)
	}
}
func setupNatsConn() {
	sc, _ = stan.Connect("test-cluster", clientID, stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
		log.Printf("Connection lost, reason: %v\n", reason)
		setupNatsConn() // Попытка переподключения
	}))
	setupSubscribe()
	log.Printf("Sucsessfully connected!")
}

func main() {
	setupNatsConn()
	select {}
}

func setOrder(jsonOrder []byte) error {
	var order models.Order
	err := json.Unmarshal(jsonOrder, &order)
	if err != nil {
		return err
	}
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	delivery := order.Delivery
	payment := order.Payment
	items := order.Items

	queryOrder := `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = db.Exec(queryOrder, order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	if (payment != models.Payment{}) {
		queryPayment := `INSERT INTO payments (order_uid, transaction, request_id, currency, provider, amount, dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
		_, err = db.Exec(queryPayment, order.OrderUid, payment.Transaction, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.Dt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
		if err != nil {
			return err
		}
	}

	if (delivery != models.Delivery{}) {
		queryPayment := `INSERT INTO deliveries (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		_, err = db.Exec(queryPayment, order.OrderUid, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
		if err != nil {
			return err
		}
	}

	for _, item := range items {
		if (item != models.Item{}) {
			queryItem := `INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9,$10,$11,$12)`
			_, err = db.Exec(queryItem, order.OrderUid, item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice, item.NMID, item.Brand, item.Status)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
