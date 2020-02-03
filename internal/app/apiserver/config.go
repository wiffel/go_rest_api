package apiserver

// Config ...
type Config struct {
	bindAddr string `toml:"bind_addr"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		bindAddr: ":8080",
	}
}
