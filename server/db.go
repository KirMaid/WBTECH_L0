package main

func main() {

}

/*func connect() {
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
		for i, _ := range columns {
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
	fmt.Println(jsonMsg)
}*/
