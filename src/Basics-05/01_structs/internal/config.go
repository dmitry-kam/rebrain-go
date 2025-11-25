package internal

import "time"

type Config struct {
	apiKey     string
	timeout    time.Duration
	maxRetries int
	private    int
}

func NewConfig() Config {
	return Config{
		apiKey:     "secret",
		timeout:    30 * time.Second,
		maxRetries: 3,
		private:    1, // guaranteed immutable
	}
}

func (c Config) APIKey() string {
	return c.apiKey
}

func (c Config) Timeout() time.Duration {
	return c.timeout
}
