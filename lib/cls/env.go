package cls

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func (c *CLS) FlushAndLoadEnv(filenames ...string) error {
	if len(filenames) != 0 {
		envMap, err := godotenv.Read(filenames...)
		if err != nil {
			log.Println("Error reading .env file:", err)
		}
		for key := range envMap {
			os.Unsetenv(key)
		}
		return godotenv.Load(filenames...)
	} else {
		envMap, err := godotenv.Read()
		if err != nil {
			log.Println("Error reading .env file:", err)
		}
		for key := range envMap {
			os.Unsetenv(key)
		}
		return godotenv.Load()
	}
}

func GetEnvString(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func GetEnvInt(key string, defaultValue int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultValue
}

func GetEnvFloat(key string, defaultValue float64) float64 {
	if val := os.Getenv(key); val != "" {
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return f
		}
	}
	return defaultValue
}
func GetEnvBool(key string, defaultValue bool) bool {
	if val := os.Getenv(key); val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defaultValue
}
