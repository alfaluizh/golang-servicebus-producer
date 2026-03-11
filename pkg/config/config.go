package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceBusConnectionString string
	ServiceBusQueue            string
}

func Load() Config {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return Config{
		ServiceBusConnectionString: os.Getenv("SERVICEBUS_CONNECTION_STRING"),
		ServiceBusQueue:            os.Getenv("SERVICEBUS_QUEUE"),
	}
}
