package database

import (
	"L0/models"
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "user=user password=password dbname=db port=5452 sslmode=disable TimeZone=Europe/Moscow"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetOrderFromDatabase(orderUid string) (models.OrderData, error) {
	db, err := ConnectDB()
	if err != nil {
		return models.OrderData{}, err
	}

	orderData, err := getOrderData(db, orderUid)

	if err != nil {
		return models.OrderData{}, err
	}

	items, err := getOrderItems(db, orderUid)

	if err != nil {
		return models.OrderData{}, err
	}

	orderData.Items = items

	return orderData, nil
}

func getOrderData(db *sql.DB, orderUid string) (models.OrderData, error) {
	rows, err := db.Query(`
		SELECT orders.order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard, 
		       deliveries.name, phone, zip, city, address, region, email, 
		       payments.transaction, request_id, currency, provider, amount, dt, bank, delivery_cost, goods_total, custom_fee
		FROM orders
		LEFT JOIN deliveries ON orders.order_uid = deliveries.order_uid
		LEFT JOIN payments ON orders.order_uid = payments.order_uid
		WHERE orders.order_uid = $1
	`, orderUid)

	if err != nil {
		return models.OrderData{}, err
	}

	defer rows.Close()

	var orderData models.OrderData

	for rows.Next() {
		var o models.Order
		var d models.Delivery
		var p models.Payment

		err := rows.Scan(
			&o.OrderUid,
			&o.TrackNumber,
			&o.Entry,
			&o.Locale,
			&o.InternalSignature,
			&o.CustomerId,
			&o.DeliveryService,
			&o.Shardkey,
			&o.SmId,
			&o.DateCreated,
			&o.OofShard,
			&d.Name,
			&d.Phone,
			&d.Zip,
			&d.City,
			&d.Address,
			&d.Region,
			&d.Email,
			&p.Transaction,
			&p.RequestId,
			&p.Currency,
			&p.Provider,
			&p.Amount,
			&p.Dt,
			&p.Bank,
			&p.DeliveryCost,
			&p.GoodsTotal,
			&p.CustomFee)

		if err != nil {
			return models.OrderData{}, err
		}

		orderData.Order = o
		orderData.Delivery = d
		orderData.Payment = p
	}

	if (orderData.Payment == models.Payment{} || (orderData.Delivery == models.Delivery{})) {
		return models.OrderData{}, sql.ErrNoRows
	}

	return orderData, err
}

func getOrderItems(db *sql.DB, orderUid string) ([]models.Item, error) {
	rows, err := db.Query(`
		SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
		FROM items
		WHERE order_uid = $1
	`, orderUid)
	if err != nil {
		return []models.Item{}, err
	}
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NMID,
			&item.Brand,
			&item.Status)

		if err != nil {
			return []models.Item{}, err
		}
		items = append(items, item)
	}
	defer rows.Close()
	return items, err
}
