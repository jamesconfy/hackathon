package logger

import (
	"log"
	"os"
)

func NewLogger() *os.File {
	f, err := os.OpenFile("./logs/gin.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		log.Panic(err)
	}

	return f
}
