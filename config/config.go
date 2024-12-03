package config

type Config struct {
	connection string
}

func NewConfig() *Config {

	return &Config{
		connection: "",
	}
}
