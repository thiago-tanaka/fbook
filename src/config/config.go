package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// APIURL is the URL of the API
	APIURL = ""

	// Port is the port that the server will listen to
	Port = 0

	// HashKey is the key used to encrypt the JWT
	HashKey []byte

	// BlockKey is the key used to encrypt the JWT
	BlockKey []byte
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
