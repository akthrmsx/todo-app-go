package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Env              string `env:"APP_ENV" envDefault:"dev"`
	Port             int    `env:"PORT" envDefault:"80"`
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"127.0.0.1"`
	DatabasePort     int    `env:"DATABASE_PORT" envDefault:"3306"`
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"todo_user"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"password"`
	DatabaseName     string `env:"DATABASE_NAME" envDefault:"todo"`
	RedisHost        string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort        int    `env:"REDIS_PORT" envDefault:"6379"`
}

func NewConfig() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
