package server

import (
	database "L0/db"
	"L0/models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
)

func setOrder(jsonOrder []byte) error {
	var orderData models.OrderData
	err := json.Unmarshal(jsonOrder, &orderData)
	if err != nil {
		return err
	}
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	order := orderData.Order
	delivery := orderData.Delivery
	payment := orderData.Payment
	items := orderData.Items

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

//func main() {
//	// Подключение к серверу NATS Streaming
//	sc, _ := stan.Connect("test-cluster", "server")
//	defer sc.Close()
//
//	//models.ConnectDatabase()
//	//connect()
//	getOrderData("some_order_uid")
//	// Бесконечный цикл для отправки сообщений
//	for {
//		err := sc.Publish("subject", []byte("Hello, World!"))
//		if err != nil {
//			log.Fatalf("Failed to publish: %v", err)
//		}
//
//		// Пауза перед следующей отправкой сообщения
//		time.Sleep(3 * time.Second)
//	}
//}

func main() {
	sc, _ := stan.Connect("test-cluster", "server")
	defer sc.Close()

	// Подписка на канал
	_, err := sc.QueueSubscribe("orders", "queue-group", func(msg *stan.Msg) {
		err := setOrder(msg.Data)
		if err != nil {
			return
		}
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	select {}
}

//func getOrder() {
//	// Подключение к серверу NATS Streaming
//	sc, _ := stan.Connect("test-cluster", "server")
//
//	getOrder("some_order_uid")
//
//	defer sc.Close()
//
//	// Подписка на канал
//	_, err := sc.QueueSubscribe("orders", "queue-group", func(msg *stan.Msg) {
//		getOrder(string(msg.Data))
//	})
//
//	if err != nil {
//		log.Fatalf("Failed to subscribe: %v", err)
//	}
//
//	// Ожидание сообщений...
//	select {}
//}

//func getOrder(orderId string) {
/*var order models.Order
var delivery models.Delivery
var payment models.Payment
var items []models.Item*/
//orders := models.Order
//models.DB.Preload("Payments").Find(&orders)
/*models.DB.Joins("left join deliveries on orders.order_uid = deliveries.order_uid").
Joins("left join payments on orders.order_uid = payments.order_uid").
Joins("left join items on orders.order_uid = items.order_uid").
Where("orders.order_uid = ?", orderId).
First(&order, &delivery, &payment, &items)*/

/*models.DB.Preload("Order").Find(&payment)
models.DB.Preload("Order").Find(&delivery)*/
//models.DB.Preload("Order").Find(&delivery)

//var order models.Order
/*var delivery Delivery
var payment Payment
var items []Item*/

/*models.DB.Preload("Deliveries").
	Preload("Payments").
	Preload("Items").
	Where("order_uid = ?", orderId).
	First(&order)

models.DB.Model(&models.Order{}).Joins(
	"LEFT JOIN deliveries ON parents.order_uid = children.order_uid AND parents.order_uid = ?",
	"some_order_uid",
)*/

/*models.DB.Joins("left join deliveries on orders.order_uid = deliveries.order_uid").
	//Joins("left join payments on orders.order_uid = payments.order_uid").
	//Joins("left join items on orders.order_uid = items.order_uid").
	Where("orders.order_uid = ?", orderId).
	First(&order)
fmt.Println(order)*/
/*orderData, err := json.Marshal(order)
if err != nil {
	log.Println(err)
	return
}
fmt.Println(string(orderData))

deliveryData, err := json.Marshal(delivery)
if err != nil {
	log.Println(err)
	return
}
fmt.Println(string(deliveryData))

paymentsData, err := json.Marshal(payment)
if err != nil {
	log.Println(err)
	return
}
fmt.Println(string(paymentsData))

itemsData, err := json.Marshal(items)
if err != nil {
	log.Println(err)
	return
}
fmt.Println(string(itemsData))*/
//}

func connect() {
	connStr := "user=user password=password dbname=db port=5452 sslmode=disable TimeZone=Europe/Moscow"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderUid := "some_order_uid"
	rows, err := db.Query("SELECT orders.*, deliveries.*, payments.*\nFROM orders\nLEFT JOIN deliveries ON orders.order_uid = deliveries.order_uid\nLEFT JOIN payments ON orders.order_uid = payments.order_uid\nWHERE orders.order_uid = $1", orderUid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	count := len(columns)

	var v struct {
		Data []interface{} `json:"data"`
	}

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}

		var m map[string]interface{}
		m = make(map[string]interface{})
		for i := range columns {
			m[columns[i]] = values[i]
		}
		v.Data = append(v.Data, m)
	}
	jsonMsg, err := json.Marshal(v)
	fmt.Println(string(jsonMsg))
}

//func getOrderData(orderUid string) {
//	connStr := "user=user password=password dbname=db port=5452 sslmode=disable TimeZone=Europe/Moscow"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	//rows, err := db.Query("SELECT orders.*\nFROM orders WHERE orders.order_uid = $1", orderUid)
//
//	rows, err := db.Query("SELECT orders.*, deliveries.*, payments.*\nFROM orders\nLEFT JOIN deliveries ON orders.order_uid = deliveries.order_uid\nLEFT JOIN payments ON orders.order_uid = payments.order_uid\nWHERE orders.order_uid = $1", orderUid)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	var orders []models2.Order
//
//	for rows.Next() {
//		var order models2.Order
//		/*err := rows.Scan(
//			&order.OrderUid,
//			&order.DateCreated,
//			&order.CustomerId,
//			&order.TrackNumber,
//			&order.Entry,
//			&order.DeliveryService,
//			&order.InternalSignature,
//			&order.Locale,
//			&order.OofShard,
//			&order.Shardkey,
//			&order.SmId)
//		if err != nil {
//			fmt.Println("Сканирование строк не удалось: %v", err)
//		}*/
//		orders = append(orders, order)
//	}
//
//	jsonMsg, err := json.Marshal(orders)
//
//	fmt.Println(string(jsonMsg))
//}
