// config.go
package main

import "github.com/gin-gonic/gin"

type ConfigMessage struct {
	LogLevel    string `json:"log_level"`
	GRPCPort    string `json:"grpc_port"`
	Environment string `json:"environment"`
	DBURL       string `json:"db_url"`
}

func ConfigHandler(c *gin.Context) {
	logLevel := getEnvVariable("LOG_LEVEL", "debug")
	grpcPort := getEnvVariable("GRPC_PORT", "80")
	environment := getEnvVariable("ENVIRONMENT", "dev")
	dbURL := getEnvVariable("DB_URL", "postgres://admin:supersecret@10.10.10.1:5432/exam-db")

	message := ConfigMessage{
		LogLevel:    logLevel,
		GRPCPort:    grpcPort,
		Environment: environment,
		DBURL:       dbURL,
	}

	c.JSON(200, message)
}
