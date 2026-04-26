package internal

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Env              string
	Port             int
	EmbeddingModel   string
	ModelPath        string
	ConnectionString string
	DefaultLogLevel  zapcore.Level
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		Env:              readString("env", "development"),
		Port:             readInt("port", 7821),
		EmbeddingModel:   readString("embedding_model", "bge-large-en-v1.5"),
		ModelPath:        readString("models", "models"),
		ConnectionString: readString("connectionString", ""),
		DefaultLogLevel:  readLogLevel("log_level"),
	}
}

func readString(key string, defaultValue string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return defaultValue
	}
	return val

}

func readInt(key string, defaultValue int) int {
	rawval := strings.TrimSpace(os.Getenv(key))
	if rawval == "" {
		return defaultValue
	}
	val, err := strconv.Atoi(strings.TrimSpace(rawval))
	if err != nil {
		return defaultValue
	}
	return val

}

func readLogLevel(key string) zapcore.Level {
	rawval := strings.TrimSpace(os.Getenv(key))
	if rawval == "" {
		return zapcore.InfoLevel
	}
	lvl, err := zapcore.ParseLevel(rawval)
	if err != nil {
		return zapcore.InfoLevel
	}
	return lvl
}
