package config

type Config struct {
    Port        int
    StoragePath string
}

func NewConfig() *Config {
    return &Config{
        Port:        5432,
        StoragePath: "data/",
    }
}
