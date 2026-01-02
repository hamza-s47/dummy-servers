package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hamza-s47/dummy-servers/handlers"
)

func main() {
	args := os.Args[1:]
	addr := ":9000"

	router := gin.Default()

	router.GET("/api/test", handlers.TestRoute)

	if len(args) > 0 && args[0] != "" {
		addr = args[0]
	}
	if err := router.Run(addr); err != nil {
		log.Fatalf("[Error] failed to start Gin server due to: %s", err.Error())
	}
}
