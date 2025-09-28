package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type (
	Environment  string
	HTTPProtocol string
)

const (
	Production  Environment = "production"
	Development Environment = "development"
	Testing     Environment = "testing"
	Staging     Environment = "staging"
	Local       Environment = "local"

	HTTPInsecure HTTPProtocol = "http"
	HTTPSecure   HTTPProtocol = "https"
)

type Config struct {
	Application
	// Database
	// Server
	Auth
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	config := Config{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return &config
}

type (
	Application struct {
		//Name        string      `env:"APP_NAME,required"`
		// Version     string      `env:"APP_VERSION,required"`
		Environment Environment `env:"APP_ENVIRONMENT,required"`
	}

	SEO struct {
		Title string `env:"SEO_TITLE"`
	}

	Server struct {
		Host     string       `env:"SERVER_HOST,required"`
		HTTPPort int          `env:"SERVER_HTTP_PORT,required" envDefault:"9090"`
		Protocol HTTPProtocol `env:"SERVER_PROTOCOL,required" envDefault:"http"`
	}

	Database struct {
		Host       string `env:"DB_HOST,required"`
		Port       int    `env:"DB_PORT,required"`
		Database   string `env:"DB_DATABASE,required"`
		User       string `env:"DB_USER,required"`
		Password   string `env:"DB_PASSWORD,required"`
		SSLMode    string `env:"DB_SSLMODE,required" envDefault:"disable"`
		Connection string `env:"DB_CONNECTION"`
	}

	Auth struct {
		Domain      string `env:"AUTH0_DOMAIN"`
		ClintID     string `env:"AUTH0_CLIENT_ID"`
		ClintSecret string `env:"AUTH0_CLIENT_SECRET"`
		CallbackURL string `env:"AUTH0_CALLBACK_URL"`
	}
)

func (d Database) ConnectionString() string {
	conn := os.Getenv("DB_CONNECTION")

	if len(conn) > 2 {
		return conn
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Database, d.SSLMode)
}
