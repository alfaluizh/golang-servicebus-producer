package main

import (
	"log"

	"github.com/alfaluizh/golang-servicebus-producer/internal/http"
	"github.com/alfaluizh/golang-servicebus-producer/internal/servicebus"
	"github.com/alfaluizh/golang-servicebus-producer/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	producer := servicebus.NewProducer(cfg.ServiceBusConnectionString, cfg.ServiceBusQueue, cfg.ServiceBusReplyQueue)

	router := gin.Default()
	handler := http.NewHandler(producer)
	router.POST("/publish", handler.Publish)
	log.Println("Server Running on :8080")

	router.Run(":8080")

}
