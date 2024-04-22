// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"./handler"
)

func getEnvVariable(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	r := gin.Default()

	r.GET("/config", handler.ConfigHandler)
	r.GET("/domain", handler.DomainHandler)
	r.GET("/load-capability", handler.LoadCapabilityHandler)

	port := ":8070"
	fmt.Printf("Server is listening on port %s...\n", port)
	r.Run(port)
}
