package main

import (
	"L0/server/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const portNumber = ":8070"

func main() {
	r := gin.Default()
	r.GET("/orders/:order_uid", getOrder)
	r.POST("/submit", func(c *gin.Context) {
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
		sendOrder(jsonData)

		// Теперь jsonData - это []byte, который вы можете использовать как захотите
		// ...

		c.JSON(http.StatusOK, gin.H{"message": "Success!"})
	})
	r.Run(":8070") // listen and serve on 0.0.0.0:8080
}

func getOrder(c *gin.Context) {
	uid := c.Param("order_uid")
	order, status := getCache(uid)
	if !status {
		var err error
		order, err = getOrderFromDatabase(uid)
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

func sendOrder(order []byte) {

}
