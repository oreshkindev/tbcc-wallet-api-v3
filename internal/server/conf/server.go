package conf

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	DB struct {
		Host         string        `default:"localhost" envconfig:"DB_HOST"`
		Name         string        `default:"postgres" envconfig:"DB_NAME"`
		User         string        `default:"postgres" envconfig:"DB_USER"`
		Password     string        `default:"root" envconfig:"DB_PASSWORD"`
		Port         int           `default:"5432" envconfig:"DB_PORT"`
		Schema       string        `default:"public" envconfig:"DB_SCHEMA"`
		PoolSize     int           `default:"5" envconfig:"DB_POOLSIZE"`
		MaxIdleConns int           `default:"3" envconfig:"DB_MAX_IDLE"`
		ConnLifetime time.Duration `default:"10m" envconfig:"DB_CONNLIFETIME"`
		Tmpl         string        `default:"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable search_path=%s application_name=%s"`
	}
}

func ParseConfig(app string) (cfg config, err error) {
	if err := envconfig.Process(app, &cfg); err != nil {
		if err := envconfig.Usage(app, &cfg); err != nil {
			return cfg, err
		}
		return cfg, err
	}
	return cfg, nil
}