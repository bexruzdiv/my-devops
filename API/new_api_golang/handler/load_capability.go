// load_capability.go
package main

import "github.com/gin-gonic/gin"

type LoadCapabilityMessage struct {
	CPU    string `json:"cpus"`
	MEMORY string `json:"memorys"`
}

func LoadCapabilityHandler(c *gin.Context) {
	cpus := getEnvVariable("CPU", "0.1")
	memorys := getEnvVariable("MEMORY", "32Mi")

	message := LoadCapabilityMessage{
		CPU:    cpus,
		MEMORY: memorys,
	}

	c.JSON(200, message)
}
