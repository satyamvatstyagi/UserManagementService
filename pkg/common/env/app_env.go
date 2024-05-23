package env

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var EnvConfig envConfigMap

type envConfigMap struct {
	DatabaseUser      string `required:"true" envconfig:"DATABASE_USER"`
	DatabasePassword  string `required:"true" envconfig:"DATABASE_PASSWORD"`
	DatabaseName      string `required:"true" envconfig:"DATABASE_NAME"`
	DatabaseHost      string `required:"true" envconfig:"DATABASE_HOST"`
	DatabasePort      string `required:"true" envconfig:"DATABASE_PORT"`
	GinMode           string `required:"true" envconfig:"GIN_MODE"`
	ServicePort       string `required:"true" envconfig:"SERVICE_PORT"`
	BasicAuthUser     string `required:"true" envconfig:"BASIC_AUTH_USER"`
	BasicAuthPassword string `required:"true" envconfig:"BASIC_AUTH_PASSWORD"`
	JWTSecretKey      string `required:"true" envconfig:"JWT_SECRET_KEY"`
	JWTExpirationTime string `required:"true" envconfig:"JWT_EXPIRATION_TIME"`
	LogFilePath       string `required:"true" envconfig:"LOG_FILE_PATH"`
}

func LoadConfig() error {

	err := envconfig.Process("", &EnvConfig)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfigWithFile(envFile string) error {

	err := godotenv.Load(envFile)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	err = envconfig.Process("", &EnvConfig)
	if err != nil {
		return err
	}
	return nil
}
