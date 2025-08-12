package config

type Config struct {
	HTTP HTTP
	Log  Log
	PG   PG
}

type HTTP struct {
	Port int `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
}

type PG struct {
	PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	URL     string `env-required:"true"                 env:"PG_URL"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	cfg.Log.Level = "debug"
	cfg.PG.URL = "postgres://postgres:postgres@localhost:5432/GoArchitectureDb?sslmode=disable"
	cfg.HTTP.Port = 8080

	return cfg, nil
}
