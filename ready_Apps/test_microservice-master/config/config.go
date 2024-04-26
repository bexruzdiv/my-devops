package config

import (
	"os"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment          string // develop, staging, production
	PostgresHost         string
	PostgresPort         int
	PostgresDatabase     string
	PostgresUser         string
	PostgresPassword     string
	LogLevel             string
	RPCPort              string
	AccountSid           string
	AuthToken            string
	PhoneNumber          string
	SMSSender            string
	PlayMobileLogin      string
	PlayMobilePassword   string
	PlayMobileUrl        string
	PlayMobileOriginator string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "my-postgres"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "deleverdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "delever"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "delever"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":2323"))

	c.AccountSid = cast.ToString(getOrReturnDefault("ACCOUNT_SID", ""))
	c.AuthToken = cast.ToString(getOrReturnDefault("AUTH_TOKEN", ""))
	c.PhoneNumber = cast.ToString(getOrReturnDefault("PHONE_NUMBER", ""))

	c.SMSSender = cast.ToString(getOrReturnDefault("", ""))

	c.PlayMobileLogin = cast.ToString(getOrReturnDefault("", ""))
	c.PlayMobilePassword = cast.ToString(getOrReturnDefault("", ""))
	c.PlayMobileUrl = cast.ToString(getOrReturnDefault("", ""))
	c.PlayMobileOriginator = cast.ToString(getOrReturnDefault("", ""))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
