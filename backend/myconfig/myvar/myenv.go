package myvar

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const (
	NAME_DB_CONNECT = "DB_CONNECT"
	NAME_PORT       = "PORT"
)

var (
	ENV_DB_CONNECT string
	ENV_PORT       string
)

func GetMyENV(EnvKey string, defaultValue ...string) string {
	value := os.Getenv(EnvKey)
	log.Info().Msgf("[ENV] %s: %s", EnvKey, value)
	if len(value) == 0 && len(defaultValue) != 0 {
		value = defaultValue[0]
	} else if len(value) == 0 && len(defaultValue) == 0 {
		value = ""
	}
	return value
}

func SetEnv() {
	godotenv.Load(".env")
	ENV_DB_CONNECT = GetMyENV(NAME_DB_CONNECT)
	ENV_PORT = GetMyENV(NAME_PORT)
}
