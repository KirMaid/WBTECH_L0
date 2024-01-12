package main

import (
	database "L0/db"
	"L0/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
)

const portNumber = ":8070"

func main() {
	r := gin.Default()
	if cacheIsEmpty() {
		err := setOrdersToCache(10)
		if err != nil {
			log.Fatalf("Failed to seed Cache: %v", err)
		}
	}
	r.GET("/orders/:order_uid", getOrder)
	r.POST("/new_order", func(c *gin.Context) {
		var formData models.Order
		if err := c.ShouldBindJSON(&formData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		jsonData, err := json.Marshal(formData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = sendOrder(jsonData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success!"})
	})
	r.Run(portNumber)
}

func getOrder(c *gin.Context) {
	uid := c.Param("order_uid")
	order, status := getCache(uid)
	if !status {
		var err error
		order, err = database.GetOrderFromDatabase(uid)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		setCache(uid, order)
	}

	jsonData, err := json.MarshalIndent(order, "", " ")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.String(http.StatusOK, "<pre>\n%s\n</pre>", string(jsonData))
}

func sendOrder(order []byte) error {
	sc, _ := stan.Connect("test-cluster", "client")
	defer sc.Close()
	err := sc.Publish("order", order)
	if err != nil {
		return err
	}
	return nil
}

func setOrdersToCache(limit int) error {
	orders, err := database.GetFirstOrdersFromDatabase(limit)
	if err != nil {
		return err
	}
	for _, order := range orders {
		setCache(order.OrderUid, order)
	}
	return nil
}
