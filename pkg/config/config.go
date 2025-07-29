package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	val, ok = os.LookupEnv(key)

	if !ok {
		log.Fatalf("env variable %q not found", key)
	}

	return val
}
