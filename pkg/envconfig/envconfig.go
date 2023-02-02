package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	DB_HOST              string   `env:"DB_HOST,notEmpty"`
	DB_USERNAME          string   `env:"DB_USERNAME,notEmpty"`
	DB_PASSWORD          string   `env:"DB_PASSWORD,notEmpty"`
	DB_NAME              string   `env:"DB_NAME,notEmpty"`
	DB_PORT              int      `env:"DB_PORT,notEmpty"`
	APP_PORT             int      `env:"APP_PORT,notEmpty"`
	ALLOWED_ORIGIN       []string `env:"ALLOWED_ORIGIN,notEmpty" envSeparator:","`
	SERVER_IP            string   `env:"SERVER_IP,notEmpty"`
	ZONE_ID              string   `env:"ZONE_ID,notEmpty"`
	CLOUDFLARE_API_TOKEN string   `env:"CLOUDFLARE_API_TOKEN,notEmpty"`
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}
