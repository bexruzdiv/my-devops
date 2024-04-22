// domain.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type DomainMessage struct {
	Endpoint string `json:"endpoint"`
	Domain   string `json:"domain"`
	LogLevel string `json:"log_level"`
}

func DomainHandler(c *gin.Context) {
	host := c.Request.Host
	fmt.Printf("Host: %s\n", host)

	logLevel := getEnvVariable("LOG_LEVEL", "debug")

	message := DomainMessage{
		Endpoint: "/domain",
		Domain:   host,
		LogLevel: logLevel,
	}

	c.JSON(200, message)
}
