package devops

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/ejson"
)

type Config struct {
	IGDB struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	} `json:"igdb"`
}

func ReadConfig() *Config {
	config := &Config{}

	secretsFilePath := os.Getenv("SECRETS_FILE_PATH")
	secretsPrivateKey := os.Getenv("SECRETS_PRIVATE_KEY")

	// do not read any secret if not defined.
	if secretsFilePath == "" && secretsPrivateKey == "" {
		panic("secrets file path and private key are not defined")
	}

	data, err := ejson.DecryptFile(secretsFilePath, "", secretsPrivateKey)
	if err != nil {
		panic(fmt.Errorf("error reading secrets file: %w", err))
	}

	if err := json.Unmarshal(data, config); err != nil {
		panic(fmt.Errorf("error unmarshalling secrets file: %w", err))
	}

	return config
}
