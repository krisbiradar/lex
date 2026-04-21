package internal

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env              string
	Port             int
	EmbeddingModel   string
	ConnectionString string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("port"))
	return &Config{
		Env:              godotenvtenvOr("env", "local"),
		Port:             port,
		EmbeddingModel:   os.Must("embeddingModel"),
		ConnectionString: os.Getenv("connectionString"),
	}
}
