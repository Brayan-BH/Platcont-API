package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetKey_PrivateJwt() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	key := os.Getenv("ENV_KEY_JWT")
	return []byte(key)
}

func GetKey_PrivateCrypto() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuración de variables de entorno")
	}
	key := os.Getenv("ENV_KEY_CRYPTO")
	return []byte(key)
}
