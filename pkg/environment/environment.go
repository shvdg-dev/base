package environment

import (
	"encoding/base64"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	isEnvLoaded bool // Used to indicate whether the .env file has been loaded.
)

// GetValueAsString retrieves an environment string value given a key.
//
// Parameters:
//   - key: The key of the environment value.
//
// Returns:
//   - The string value of the environment key.
func GetValueAsString(key string) string {
	if !isEnvLoaded {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		isEnvLoaded = true
	}
	return os.Getenv(key)
}

// GetValueAsBytes retrieves an environment []byte value given a key.
//
// Parameters:
//   - key: The key of the environment value.
//
// Returns:
//   - []byte: The []byte value of the environment key.
//   - error: An error if any occurred during the decoding.
func GetValueAsBytes(key string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(GetValueAsString(key))
}
