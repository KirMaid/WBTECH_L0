package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
)

const portNumber = ":8080"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We're live !!!"))
	// Подключение к серверу NATS Streaming
	sc, _ := stan.Connect("test-cluster", "client")
	defer sc.Close()

	// Подписка на канал
	_, err := sc.QueueSubscribe("subject", "queue-group", func(msg *stan.Msg) {
		// Обработка полученного сообщения
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	// Ожидание сообщений...
	select {}
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("client/views/*")
	router.POST("/submit", func(c *gin.Context) {
		var text string
		// Валидация входных данных и связывание их с переменной text
		if err := c.ShouldBindQuery(&text); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Используйте данные формы
		c.JSON(http.StatusOK, gin.H{"text": text})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8070")
}

func getOrderId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	orderId := r.FormValue("orderId")
	getOrder(orderId)
}

func getOrder(orderId string) {
	sendOrderId(orderId)
	reciveOrder()
	/*// Подключение к NATS Streaming Server
	sc, _ := stan.Connect("test-cluster", "client")
	defer sc.Close()

	// Отправка сообщения
	err := sc.Publish("orderIds", []byte(orderId))
	if err != nil {
		log.Fatal(err)
	}*/
}

func sendOrderId(orderId string) {
	// Подключение к NATS Streaming Server
	sc, _ := stan.Connect("test-cluster", "client")
	defer sc.Close()

	// Отправка сообщения
	err := sc.Publish("orderIds", []byte(orderId))
	if err != nil {
		log.Fatal(err)
	}
}

func reciveOrder() {
	// Подключение к серверу NATS Streaming
	sc, _ := stan.Connect("test-cluster", "client")
	defer sc.Close()

	// Подписка на канал
	_, err := sc.QueueSubscribe("orders", "queue-group", func(msg *stan.Msg) {
		// Обработка полученного сообщения
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	// Ожидание сообщений...
	select {}
}

//	http.HandleFunc("/", HomePageHandler)
//	http.ListenAndServe(portNumber, nil)
//}

//func setupRouter() *gin.Engine {
//	router := gin.Default()
//
//	router.StaticFS("/static", http.Dir("./public/static"))
//
//	return router
//}
