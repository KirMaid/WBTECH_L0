package main

import (
	database "L0/db"
	"L0/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"net/http"
)

const portNumber = ":8070"

func main() {
	r := gin.Default()
	r.GET("/orders/:order_uid", getOrder)
	r.POST("/new_order", func(c *gin.Context) {
		var formData models.OrderData

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
	r.Run(portNumber) // listen and serve on 0.0.0.0:8080
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
