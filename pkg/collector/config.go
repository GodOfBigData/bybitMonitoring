package collector

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigExchange struct {
	apiKey      string
	apiSecret   string
	url         string
	recv_window string
	address     string
	topic       string
	symbol      string
}

func (c *ConfigExchange) GetApiKey() string {
	return c.apiKey
}

func (c *ConfigExchange) GetSecret() string {
	return c.apiSecret
}

func (c *ConfigExchange) GetUrl() string {
	return c.url
}

func (c *ConfigExchange) GetRecvWindow() string {
	return c.recv_window
}

func (c *ConfigExchange) GetAddress() string {
	return c.address
}

func (c *ConfigExchange) GetTopic() string {
	return c.topic
}

func (c *ConfigExchange) GetSymbol() string {
	return c.symbol
}

type Config struct {
	ConfigExchange ConfigExchange
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func New(pathEnv string) (*Config, error) {
	if err := godotenv.Load(pathEnv); err != nil {
		return &Config{}, err
	}
	return &Config{
		ConfigExchange: ConfigExchange{
			apiKey:      getEnv("api_key", ""),
			apiSecret:   getEnv("api_secret", ""),
			url:         getEnv("url", ""),
			recv_window: getEnv("recv_window", ""),
			address:     getEnv("address", ""),
			topic:       getEnv("topic", ""),
			symbol:      getEnv("symbol", ""),
		},
	}, nil
}
